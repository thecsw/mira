package mira

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Reddit) MiraRequest(method string, target string, payload map[string]string) ([]byte, error) {
	values := "?"
	for i, v := range payload {
		v = url.QueryEscape(v)
		values += fmt.Sprintf("%s=%s&", i, v)
	}
	values = values[:len(values)-1]
	r, err := http.NewRequest(method, target+values, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "Bearer "+c.Token)
	response, err := c.Client.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	data := buf.Bytes()
	if err := findRedditError(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (c *Reddit) Me() *Reddit {
	c.Chain.Type = "me"
	return c
}

func (c *Reddit) Subreddit(name string) *Reddit {
	c.Chain.Name = name
	c.Chain.Type = "subreddit"
	return c
}

func (c *Reddit) Submission(name string) *Reddit {
	c.Chain.Name = name
	c.Chain.Type = "submission"
	return c
}

func (c *Reddit) Comment(name string) *Reddit {
	c.Chain.Name = name
	c.Chain.Type = "comment"
	return c
}

func (c *Reddit) Redditor(name string) *Reddit {
	c.Chain.Name = name
	c.Chain.Type = "redditor"
	return c
}

func (c *Reddit) Submissions(sort string, tdur string, limit int) ([]PostListingChild, error) {
	switch c.Chain.Type {
	case "subreddit":
		return c.getSubredditPosts(c.Chain.Name, sort, tdur, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for submissions", c.Chain.Type)
	}
}

func (c *Reddit) SubmissionsAfter(last string, limit int) ([]PostListingChild, error) {
	switch c.Chain.Type {
	case "subreddit":
		return c.getSubredditPostsAfter(c.Chain.Name, last, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for submissions", c.Chain.Type)
	}
}

func (c *Reddit) Comments(sort string, tdur string, limit int) ([]Comment, error) {
	switch c.Chain.Type {
	case "subreddit":
		return c.getSubredditComments(c.Chain.Name, sort, tdur, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for submissions", c.Chain.Type)
	}
}

func (c *Reddit) Info() (interface{}, error) {
	switch c.Chain.Type {
	case "me":
		return c.getMe()
	case "submission":
		return c.getSubmission(c.Chain.Name)
	case "comment":
		return c.getComment(c.Chain.Name)
	case "subreddit":
		return c.getSubreddit(c.Chain.Name)
	case "redditor":
		return c.getUser(c.Chain.Name)
	default:
		return nil, fmt.Errorf("returning type is not defined")
	}
}

func (c *Reddit) getMe() (*Me, error) {
	target := RedditOauth + "/api/v1/me"
	ans, err := c.MiraRequest("GET", target, nil)
	if err != nil {
		return nil, err
	}
	ret := &Me{}
	json.Unmarshal(ans, ret)
	return ret, nil
}

func (c *Reddit) getSubmission(id string) (*PostListing, error) {
	target := RedditOauth + "/api/info.json"
	ans, err := c.MiraRequest("GET", target, map[string]string{
		"id": id,
	})
	ret := &PostListing{}
	json.Unmarshal(ans, ret)
	return ret, err
}

func (c *Reddit) getComment(id string) (*Comment, error) {
	target := RedditOauth + "/api/info.json"
	ans, err := c.MiraRequest("GET", target, map[string]string{
		"id": id,
	})
	ret := CommentListing{}
	json.Unmarshal(ans, ret)
	return &ret.GetChildren()[0], err
}

// This function will return the submission id of a comment
//
// Comment id has form of t1_... where submission is prefixed with t3_...
//
// Comment structures in themselves do not have submission id included,
// they only have a parent_id field that points to a parent comment or a
// submission. If it does not point directly at the submission, we need to
// make iterative calls until we bump into an id that fits the submission
// prefix. It may take several calls.
//
// For example:
//
// - If comment is first-level, we make one call to get the object and
// extract the submission id. If you already have a Go struct at hand,
// please just invoke .GetParentId() to get the parent id
//
// - If comment is second-level, it would take two calls to extact the
// submission id. If you want to save a call, you can pass a parent_id
// instead that would take 1 call instead of 2.
//
// - If comment is N-level, it would take N calls. If you aleady have an
// object, just pass in its object, so it would take N-1 calls.
//
// NOTE: If any error occurs, the method will return on error object.
// If it takes more than 12 calls, the function bails out.
func (c *Reddit) GetSubmissionFromComment(comment_id string) (string, error) {
	current := comment_id
	// Not a comment passed
	if string(current[1]) != "1" {
		return "", errors.New("the passed ID is not a comment")
	}
	target := RedditOauth + "/api/info.json"
	temp := CommentListing{}
	tries := 0
	for string(current[1]) != "3" {
		ans, err := c.MiraRequest("GET", target, map[string]string{
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
			return "", errors.New(fmt.Sprintf("Exceeded the maximum number of iterations: %v", c.Values.GetSubmissionFromCommentTries))
		}
	}
	return current, nil
}

func (c *Reddit) getUser(name string) (*Redditor, error) {
	target := RedditOauth + "/user/" + name + "/about"
	ans, err := c.MiraRequest("GET", target, nil)
	ret := &Redditor{}
	json.Unmarshal(ans, ret)
	return ret, err
}

func (c *Reddit) getSubreddit(name string) (*Subreddit, error) {
	target := RedditOauth + "/r/" + name + "/about"
	ans, err := c.MiraRequest("GET", target, nil)
	ret := &Subreddit{}
	json.Unmarshal(ans, ret)
	return ret, err
}

// Get submisssions from a subreddit up to a specified limit sorted by the given parameter
//
// Sorting options: "hot", "new", "top", "rising", "controversial", "random"
//
// Time options: "all", "year", "month", "week", "day", "hour"
//
// Limit is any numerical value, so 0 <= limit <= 100
func (c *Reddit) getSubredditPosts(sr string, sort string, tdur string, limit int) ([]PostListingChild, error) {
	target := RedditOauth + "/r/" + sr + "/" + sort + ".json"
	ans, err := c.MiraRequest("GET", target, map[string]string{
		"limit": strconv.Itoa(limit),
		"t":     tdur,
	})
	ret := &PostListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getSubredditComments(sr string, sort string, tdur string, limit int) ([]Comment, error) {
	target := RedditOauth + "/r/" + sr + "/comments.json"
	ans, err := c.MiraRequest("GET", target, map[string]string{
		"sort":  sort,
		"limit": strconv.Itoa(limit),
		"t":     tdur,
	})
	ret := &CommentListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) GetSubmissionComments(sr string, post_id string, sort string, limit int) ([]Comment, []string, error) {
	if string(post_id[1]) != "3" {
		return nil, nil, errors.New("the passed ID36 is not a submission")
	}
	target := RedditOauth + "/r/" + sr + "/comments/" + post_id[3:]
	ans, err := c.MiraRequest("GET", target, map[string]string{
		"sort":     sort,
		"limit":    strconv.Itoa(limit),
		"showmare": strconv.FormatBool(false),
	})
	if err != nil {
		return nil, nil, err
	}
	temp := make([]CommentListing, 0, 8)
	json.Unmarshal(ans, &temp)
	ret := make([]Comment, 0, 8)
	for _, v := range temp {
		comments := v.GetChildren()
		for _, v2 := range comments {
			ret = append(ret, v2)
		}
	}
	// Cut off the "more" kind
	children := ret[len(ret)-1].Data.Children
	ret = ret[:len(ret)-1]
	return ret, children, nil
}

// Get submisssions from a subreddit up to a specified limit sorted by the given parameters
// and with specified anchor
//
// Sorting options: "hot", "new", "top", "rising", "controversial", "random"
//
// Limit is any numerical value, so 0 <= limit <= 100
//
// Anchor options are submissions full thing, for example: t3_bqqwm3
func (c *Reddit) getSubredditPostsAfter(sr string, last string, limit int) ([]PostListingChild, error) {
	target := RedditOauth + "/r/" + sr + "/new.json"
	ans, err := c.MiraRequest("GET", target, map[string]string{
		"limit":  strconv.Itoa(limit),
		"before": last,
	})
	ret := &PostListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) CommentsAfter(sort string, last string, limit int) ([]Comment, error) {
	switch c.Chain.Type {
	case "subreddit":
		return c.getSubredditCommentsAfter(c.Chain.Name, sort, last, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for comments", c.Chain.Type)
	}
}

func (c *Reddit) getSubredditCommentsAfter(sr string, sort string, last string, limit int) ([]Comment, error) {
	target := RedditOauth + "/r/" + sr + "/comments.json"
	ans, err := c.MiraRequest("GET", target, map[string]string{
		"sort":   sort,
		"limit":  strconv.Itoa(limit),
		"before": last,
	})
	ret := &CommentListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) Submit(title string, text string) (*Submission, error) {
	err := c.checkType("subreddit")
	if err != nil {
		return nil, err
	}
	target := RedditOauth + "/api/submit"
	ans, err := c.MiraRequest("POST", target, map[string]string{
		"title":    title,
		"sr":       c.Chain.Name,
		"text":     text,
		"kind":     "self",
		"resubmit": "true",
		"api_type": "json",
	})
	ret := &Submission{}
	json.Unmarshal(ans, ret)
	return ret, err
}

func (c *Reddit) Reply(text string) (*CommentWrap, error) {
	err := c.checkType("comment")
	if err != nil {
		return nil, err
	}

	target := RedditOauth + "/api/comment"
	ans, err := c.MiraRequest("POST", target, map[string]string{
		"text":     text,
		"thing_id": c.Chain.Name,
		"api_type": "json",
	})
	ret := &CommentWrap{}
	json.Unmarshal(ans, ret)
	return ret, err
}

func (c *Reddit) Save(text string) (*CommentWrap, error) {
	err := c.checkType("submission")
	if err != nil {
		return nil, err
	}
	target := RedditOauth + "/api/comment"
	ans, err := c.MiraRequest("POST", target, map[string]string{
		"text":     text,
		"thing_id": c.Chain.Name,
		"api_type": "json",
	})
	ret := &CommentWrap{}
	json.Unmarshal(ans, ret)
	return ret, err
}

func (c *Reddit) Delete() error {
	err := c.checkType("comment", "submission")
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/del"
	_, err = c.MiraRequest("POST", target, map[string]string{
		"id":       c.Chain.Name,
		"api_type": "json",
	})
	return err
}

func (c *Reddit) Approve() error {
	err := c.checkType("comment")
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/approve"
	_, err = c.MiraRequest("POST", target, map[string]string{
		"id":       c.Chain.Name,
		"api_type": "json",
	})
	return err
}

func (c *Reddit) Distinguish(how string, sticky bool) error {
	err := c.checkType("comment")
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/distinguish"
	_, err = c.MiraRequest("POST", target, map[string]string{
		"id":       c.Chain.Name,
		"how":      how,
		"sticky":   strconv.FormatBool(sticky),
		"api_type": "json",
	})
	return err
}

func (c *Reddit) Edit(text string) (*CommentWrap, error) {
	err := c.checkType("comment")
	if err != nil {
		return nil, err
	}
	target := RedditOauth + "/api/editusertext"
	ans, err := c.MiraRequest("POST", target, map[string]string{
		"text":     text,
		"thing_id": c.Chain.Name,
		"api_type": "json",
	})
	ret := &CommentWrap{}
	json.Unmarshal(ans, ret)
	return ret, err
}

func (c *Reddit) Compose(subject, text string) error {
	err := c.checkType("redditor")
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/compose"
	_, err = c.MiraRequest("POST", target, map[string]string{
		"subject":  subject,
		"text":     text,
		"to":       c.Chain.Name,
		"api_type": "json",
	})
	return err
}

func (c *Reddit) ReadMessage(message_id string) error {
	target := RedditOauth + "/api/read_message"
	_, err := c.MiraRequest("POST", target, map[string]string{
		"id": message_id,
	})
	return err
}

func (c *Reddit) ReadAllMessages() error {
	target := RedditOauth + "/api/read_all_messages"
	_, err := c.MiraRequest("POST", target, nil)
	return err
}

func (c *Reddit) ListUnreadMessages() ([]Comment, error) {
	target := RedditOauth + "/message/unread"
	ans, err := c.MiraRequest("GET", target, map[string]string{
		"mark": "true",
	})
	ret := &CommentListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) SubredditUpdateSidebar(sr, text string) ([]byte, error) {
	target := RedditOauth + "/api/site_admin"
	return c.MiraRequest("POST", target, map[string]string{
		"sr":          sr,
		"name":        "None",
		"description": text,
		"title":       sr,
		"wikimode":    "anyone",
		"link_type":   "any",
		"type":        "public",
		"api_type":    "json",
	})
}

func (c *Reddit) checkType(rtype ...string) error {
	if c.Chain.Name == "" {
		return fmt.Errorf("identifier is empty")
	}
	if !findElem(c.Chain.Type, rtype) {
		return fmt.Errorf("the passed type is not a valid type for this call | expected: %s", strings.Join(rtype, ", "))
	}
	return nil
}

func findElem(elem string, arr []string) bool {
	for _, v := range arr {
		if elem == v {
			return true
		}
	}
	return false
}

type RedditErr struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func findRedditError(data []byte) error {
	object := &RedditErr{}
	json.Unmarshal(data, object)
	if object.Message != "" || object.Error != "" {
		return fmt.Errorf("%s | error code: %s", object.Message, object.Error)
	}
	return nil
}
