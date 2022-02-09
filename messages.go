package mira

import (
	"encoding/json"
	"net/http"

	"github.com/thecsw/mira/models"
)

// Compose will send a message to next redditor.
func (c *Reddit) Compose(subject, text string) error {
	name, _, err := c.checkType(redditorType)
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/compose"
	_, err = c.MiraRequest(http.MethodPost, target, map[string]string{
		"subject":  subject,
		"text":     text,
		"to":       name,
		"api_type": JsonAPI,
	})
	return err
}

// ReadMessage marks the next comment/message as read.
func (c *Reddit) ReadMessage(messageID string) error {
	_, _, err := c.checkType(meType)
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/read_message"
	_, err = c.MiraRequest(http.MethodPost, target, map[string]string{
		"id": messageID,
	})
	return err
}

// ReadAllMessages uses ReadMessage on all unread messages.
func (c *Reddit) ReadAllMessages() error {
	_, _, err := c.checkType(meType)
	if err != nil {
		return err
	}
	target := RedditOauth + "/api/read_all_messages"
	_, err = c.MiraRequest(http.MethodPost, target, nil)
	return err
}

// ListUnreadMessages returns a list of all unread messages.
func (c *Reddit) ListUnreadMessages() ([]models.Comment, error) {
	_, _, err := c.checkType(meType)
	if err != nil {
		return nil, err
	}
	target := RedditOauth + "/message/unread"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"mark": "false",
	})
	ret := &models.CommentListing{}
	json.Unmarshal(ans, ret)
	return ret.GetChildren(), err
}
