package mira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/thecsw/mira/v4/models"
)

// ModQueue returns modqueue entries from a subreddit up to a specified limit sorted by the given parameters
// Limit is any numerical value, so 0 <= limit <= 100.
func (c *Reddit) ModQueue(limit int) ([]models.ModQueueListingChild, error) {
	name, ttype := c.getQueue()
	switch ttype {
	case subredditType:
		return c.getSubredditModQueue(name, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for modqueue", ttype)
	}
}

// ModQueueAfter returns new modqueue entries from a subreddit
//
// # Last is the anchor of a modqueue entry id
//
// Limit is any numerical value, so 0 <= limit <= 100.
func (c *Reddit) ModQueueAfter(last string, limit int) ([]models.ModQueueListingChild, error) {
	name, ttype := c.getQueue()
	switch ttype {
	case subredditType:
		return c.getSubredditModQueueAfter(name, last, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for modqueue", ttype)
	}
}

func (c *Reddit) getSubredditModQueue(sr string, limit int) ([]models.ModQueueListingChild, error) {
	target := RedditOauth + "/r/" + sr + "/about/modqueue.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"limit": strconv.Itoa(limit),
	})
	if err != nil {
		return nil, errors.Wrap(err, "mira request failed in getSubredditModQueue")
	}
	ret := models.ModQueueListing{}
	err = json.Unmarshal(ans, &ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getSubredditModQueueAfter(sr string, last string, limit int) ([]models.ModQueueListingChild, error) {
	target := RedditOauth + "/r/" + sr + "/about/modqueue.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"limit":  strconv.Itoa(limit),
		"before": last,
	})
	if err != nil {
		return nil, errors.Wrap(err, "mira request failed in getSubredditModQueueAfter")
	}
	ret := models.ModQueueListing{}
	err = json.Unmarshal(ans, &ret)
	return ret.GetChildren(), err
}
