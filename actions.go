package mira

import (
	"encoding/json"
	"net/http"

	"github.com/thecsw/mira/models"
)

// Submit submits a submission to a subreddit.
func (c *Reddit) Submit(title string, text string) (models.Submission, error) {
	ret := &models.Submission{}
	name, _, err := c.checkType(subredditType)
	if err != nil {
		return *ret, err
	}
	target := RedditOauth + "/api/submit"
	ans, err := c.MiraRequest(http.MethodPost, target, map[string]string{
		"title":    title,
		"sr":       name,
		"text":     text,
		"kind":     "self",
		"resubmit": "true",
		"api_type": "json",
	})
	json.Unmarshal(ans, ret)
	return *ret, err
}

// Reply replies to a comment with text.
func (c *Reddit) Reply(text string) (models.CommentWrap, error) {
	ret := &models.CommentWrap{}
	name, _, err := c.checkType(commentType)
	if err != nil {
		return *ret, err
	}
	target := RedditOauth + "/api/comment"
	ans, err := c.MiraRequest(http.MethodPost, target, map[string]string{
		"text":     text,
		"thing_id": name,
		"api_type": "json",
	})
	json.Unmarshal(ans, ret)
	return *ret, err
}

// ReplyWithID is the same as Reply but with explicit passing comment id.
func (c *Reddit) ReplyWithID(name, text string) (models.CommentWrap, error) {
	ret := &models.CommentWrap{}
	target := RedditOauth + "/api/comment"
	ans, err := c.MiraRequest(http.MethodPost, target, map[string]string{
		"text":     text,
		"thing_id": name,
		"api_type": "json",
	})
	json.Unmarshal(ans, ret)
	return *ret, err
}

// Save posts a comment to a submission.
func (c *Reddit) Save(text string) (models.CommentWrap, error) {
	ret := &models.CommentWrap{}
	name, _, err := c.checkType(submissionType)
	if err != nil {
		return *ret, err
	}
	target := RedditOauth + "/api/comment"
	ans, err := c.MiraRequest(http.MethodPost, target, map[string]string{
		"text":     text,
		"thing_id": name,
		"api_type": "json",
	})
	json.Unmarshal(ans, ret)
	return *ret, err
}

// SaveWithID is the same as Save but with explicitely passing.
func (c *Reddit) SaveWithID(name, text string) (models.CommentWrap, error) {
	ret := &models.CommentWrap{}
	target := RedditOauth + "/api/comment"
	ans, err := c.MiraRequest(http.MethodPost, target, map[string]string{
		"text":     text,
		"thing_id": name,
		"api_type": "json",
	})
	json.Unmarshal(ans, ret)
	return *ret, err
}

// Delete deletes whatever is next in the queue.
func (c *Reddit) Delete() error {
	name, _, err := c.checkType(commentType, submissionType)
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/del"
	_, err = c.MiraRequest(http.MethodPost, target, map[string]string{
		"id":       name,
		"api_type": "json",
	})
	return err
}

// Edit will edit the next queued comment.
func (c *Reddit) Edit(text string) (models.CommentWrap, error) {
	ret := &models.CommentWrap{}
	name, _, err := c.checkType(commentType, submissionType)
	if err != nil {
		return *ret, err
	}
	target := RedditOauth + "/api/editusertext"
	ans, err := c.MiraRequest(http.MethodPost, target, map[string]string{
		"text":     text,
		"thing_id": name,
		"api_type": "json",
	})
	json.Unmarshal(ans, ret)
	return *ret, err
}
