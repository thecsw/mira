package mira

import (
	"net/http"
	"strconv"
)

// Approve is a mod tool to approve a comment or a submission
// Will fail if not a mod.
func (c *Reddit) Approve() error {
	name, _, err := c.checkType(commentType)
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/approve"
	_, err = c.MiraRequest(http.MethodPost, target, map[string]string{
		"id":       name,
		"api_type": "json",
	})
	return err
}

// Distinguish is a mod tool to distinguish a comment or a submission
// Will fail if not a mod.
func (c *Reddit) Distinguish(how string, sticky bool) error {
	name, _, err := c.checkType(commentType)
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/distinguish"
	_, err = c.MiraRequest(http.MethodPost, target, map[string]string{
		"id":       name,
		"how":      how,
		"sticky":   strconv.FormatBool(sticky),
		"api_type": "json",
	})
	return err
}

// UpdateSidebar updates subreddit's sidebar, Needs mod privileges.
func (c *Reddit) UpdateSidebar(text string) error {
	name, _, err := c.checkType(subredditType)
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/site_admin"
	_, err = c.MiraRequest(http.MethodPost, target, map[string]string{
		"sr":          name,
		"name":        "None",
		"description": text,
		"title":       name,
		"wikimode":    "anyone",
		"link_type":   "any",
		"type":        "public",
		"api_type":    "json",
	})
	return err
}

// SelectFlair sets a submission flair.
func (c *Reddit) SelectFlair(text string) error {
	name, _, err := c.checkType(submissionType)
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/selectflair"
	_, err = c.MiraRequest(http.MethodPost, target, map[string]string{
		"link":     name,
		"text":     text,
		"api_type": "json",
	})
	return err
}

// SelectFlairWithID sets submission flair with explicit ID.
func (c *Reddit) SelectFlairWithID(name, text string) error {
	target := RedditOauth + "/api/selectflair"
	_, err := c.MiraRequest(http.MethodPost, target, map[string]string{
		"link":     name,
		"text":     text,
		"api_type": "json",
	})
	return err
}

// UserFlair updates user's flair in a sub. Needs mod permissions.
func (c *Reddit) UserFlair(user, text string) error {
	name, _, err := c.checkType(subredditType)
	if err != nil {
		return err
	}
	target := RedditOauth + "/r/" + name + "/api/flair"
	_, err = c.MiraRequest(http.MethodPost, target, map[string]string{
		"name":     user,
		"text":     text,
		"api_type": "json",
	})
	return err
}

// UserFlairWithID is the same as UserFlair but explicit redditor name.
func (c *Reddit) UserFlairWithID(name, user, text string) error {
	target := RedditOauth + "/r/" + name + "/api/flair"
	_, err := c.MiraRequest(http.MethodPost, target, map[string]string{
		"name":     user,
		"text":     text,
		"api_type": "json",
	})
	return err
}
