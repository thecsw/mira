package mira

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/thecsw/mira/models"
)

// Comments returns comments from a subreddit up to a specified limit sorted by the given parameters
//
// Sorting options: "hot", "new", "top", "rising", "controversial", "random"
//
// Duration options: "hour", "day", "week", "year", "all"
//
// Limit is any numerical value, so 0 <= limit <= 100.
func (c *Reddit) Comments(sort string, tdur string, limit int) ([]models.Comment, error) {
	name, ttype := c.getQueue()
	switch ttype {
	case subredditType:
		return c.getSubredditComments(name, sort, tdur, limit)
	case submissionType:
		comments, _, err := c.getSubmissionComments(name, sort, tdur, limit)
		if err != nil {
			return nil, err
		}
		return comments, nil
	case redditorType:
		return c.getRedditorComments(name, sort, tdur, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for comments", ttype)
	}
}

// CommentsAfter returns new comments from a subreddit
//
// Last is the anchor of a comment id
//
// Limit is any numerical value, so 0 <= limit <= 100.
func (c *Reddit) CommentsAfter(sort string, last string, limit int) ([]models.Comment, error) {
	name, ttype := c.getQueue()
	switch ttype {
	case subredditType:
		return c.getSubredditCommentsAfter(name, sort, last, limit)
	case redditorType:
		return c.getRedditorCommentsAfter(name, sort, last, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for comments", ttype)
	}
}

func (c *Reddit) getComment(id string) (models.Comment, error) {
	target := RedditOauth + "/api/info.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"id": id,
	})
	ret := &models.CommentListing{}
	json.Unmarshal(ans, ret)
	if len(ret.GetChildren()) < 1 {
		return models.Comment{}, fmt.Errorf("id not found")
	}
	return ret.GetChildren()[0], err
}

func (c *Reddit) getSubredditComments(sr string, sort string, tdur string, limit int) ([]models.Comment, error) {
	target := RedditOauth + "/r/" + sr + "/comments.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"sort":  sort,
		"limit": strconv.Itoa(limit),
		"t":     tdur,
	})
	ret := &models.CommentListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getSubredditCommentsAfter(sr string, sort string, last string, limit int) ([]models.Comment, error) {
	target := RedditOauth + "/r/" + sr + "/comments.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"sort":   sort,
		"limit":  strconv.Itoa(limit),
		"before": last,
	})
	ret := &models.CommentListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getRedditorComments(user string, sort string, tdur string, limit int) ([]models.Comment, error) {
	target := RedditOauth + "/u/" + user + "/comments.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"sort":  sort,
		"limit": strconv.Itoa(limit),
		"t":     tdur,
	})
	ret := &models.CommentListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getRedditorCommentsAfter(user string, sort string, last string, limit int) ([]models.Comment, error) {
	target := RedditOauth + "/u/" + user + "/comments.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"sort":   sort,
		"limit":  strconv.Itoa(limit),
		"before": last,
	})
	ret := &models.CommentListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getSubmissionComments(postID string, sort string, tdur string, limit int) ([]models.Comment, []string, error) {
	if string(postID[1]) != "3" {
		return nil, nil, errors.New("the passed ID36 is not a submission")
	}
	target := RedditOauth + "/comments/" + postID[3:]
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"sort":     sort,
		"limit":    strconv.Itoa(limit),
		"showmore": strconv.FormatBool(true),
		"t":        tdur,
	})
	if err != nil {
		return nil, nil, err
	}
	temp := make([]models.CommentListing, 0, 8)
	json.Unmarshal(ans, &temp)
	ret := make([]models.Comment, 0, 8)
	for _, v := range temp {
		comments := v.GetChildren()
		ret = append(ret, comments...)
	}
	// Cut off the "more" kind
	children := ret[len(ret)-1].Data.Children
	ret = ret[:len(ret)-1]
	return ret, children, nil
}
