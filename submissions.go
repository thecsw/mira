package mira

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/thecsw/mira/v4/models"
)

// Submissions returns submissions from a subreddit up to a specified limit sorted by the given parameters
//
// Sorting options: `Hot`, `New`, `Top`, `Rising`, `Controversial`, `Random`
//
// Duration options: `Hour`, `Day`, `Week`, `Year`, `All`
//
// Limit is any numerical value, so 0 <= limit <= 100.
func (c *Reddit) Submissions(sort string, tdur string, limit int) ([]models.PostListingChild, error) {
	name, ttype := c.getQueue()
	switch ttype {
	case subredditType:
		return c.getSubredditPosts(name, sort, tdur, limit)
	case redditorType:
		return c.getRedditorPosts(name, sort, tdur, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for submissions", ttype)
	}
}

// SubmissionsAfter returns new submissions from a subreddit
//
// # Last is the anchor of a submission id
//
// Limit is any numerical value, so 0 <= limit <= 100.
func (c *Reddit) SubmissionsAfter(last string, limit int) ([]models.PostListingChild, error) {
	name, ttype := c.getQueue()
	switch ttype {
	case subredditType:
		return c.getSubredditPostsAfter(name, last, limit)
	case redditorType:
		return c.getRedditorPostsAfter(name, last, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for submissions", ttype)
	}
}

// ExtractSubmission extracts submission id from last pushed object
// does not make an api call like .Root(), use this instead.
func (c *Reddit) ExtractSubmission() (string, error) {
	name, _, err := c.checkType(commentType)
	if err != nil {
		return "", err
	}
	info, err := c.Comment(name).Info()
	if err != nil {
		return "", err
	}
	link := info.GetUrl()
	reg := regexp.MustCompile(`comments/([^/]+)/`)
	res := reg.FindStringSubmatch(link)
	if len(res) < 1 {
		return "", errors.New("couldn't extract submission id")
	}
	return "t3_" + res[1], nil
}

// Root will return the submission id of a comment
// Very expensive on API calls, please use .ExtractSubmission() instead.
func (c *Reddit) Root() (string, error) {
	name, _, err := c.checkType(commentType)
	if err != nil {
		return "", err
	}
	current := name
	// Not a comment passed
	if string(current[1]) != "1" {
		return "", errors.New("the passed ID is not a comment")
	}
	target := RedditOauth + "/api/info.json"
	temp := models.CommentListing{}
	tries := 0
	for string(current[1]) != "3" {
		ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
			"id": current,
		})
		if err != nil {
			return "", err
		}
		json.Unmarshal(ans, &temp)
		if len(temp.Data.Children) < 1 {
			return "", errors.New("could not find the requested comment")
		}
		current = temp.Data.Children[0].GetParentId()
		tries++
		if tries > c.Values.GetSubmissionFromCommentTries {
			return "", fmt.Errorf("exceeded the maximum number of iterations: %v",
				c.Values.GetSubmissionFromCommentTries)
		}
	}
	return current, nil
}

func (c *Reddit) getRedditorPosts(user string, sort string, tdur string, limit int) ([]models.PostListingChild, error) {
	target := RedditOauth + "/u/" + user + "/submitted/" + sort + ".json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"limit": strconv.Itoa(limit),
		"t":     tdur,
	})
	ret := &models.PostListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getRedditorPostsAfter(user string, last string, limit int) ([]models.PostListingChild, error) {
	target := RedditOauth + "/u/" + user + "/submitted/new.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"limit":  strconv.Itoa(limit),
		"before": last,
	})
	ret := &models.PostListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getSubredditPosts(sr string, sort string, tdur string, limit int) ([]models.PostListingChild, error) {
	target := RedditOauth + "/r/" + sr + "/" + sort + ".json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"limit": strconv.Itoa(limit),
		"t":     tdur,
	})
	ret := &models.PostListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getSubredditPostsAfter(sr string, last string, limit int) ([]models.PostListingChild, error) {
	target := RedditOauth + "/r/" + sr + "/new.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"limit":  strconv.Itoa(limit),
		"before": last,
	})
	ret := &models.PostListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}
