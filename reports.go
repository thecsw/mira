package mira

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/thecsw/mira/models"
)

// Reports returns report entries from a subreddit up to a specified limit sorted by the given parameters
// Limit is any numerical value, so 0 <= limit <= 100.
func (c *Reddit) Reports(limit int) ([]models.ReportListingChild, error) {
	name, ttype := c.getQueue()
	switch ttype {
	case subredditType:
		return c.getSubredditReports(name, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for reports", ttype)
	}
}

// ReportsAfter returns new report entries from a subreddit
//
// Last is the anchor of a modqueue entry id
//
// Limit is any numerical value, so 0 <= limit <= 100.
func (c *Reddit) ReportsAfter(last string, limit int) ([]models.ReportListingChild, error) {
	name, ttype := c.getQueue()
	switch ttype {
	case subredditType:
		return c.getSubredditReportsAfter(name, last, limit)
	default:
		return nil, fmt.Errorf("'%s' type does not have an option for reports", ttype)
	}
}

func unMarshalReports(ans []byte, mql models.ReportListing) (models.ReportListing, error) {
	json.Unmarshal(ans, &mql)
	return mql, nil
}

func (c *Reddit) getSubredditReports(sr string, limit int) ([]models.ReportListingChild, error) {
	target := RedditOauth + "/r/" + sr + "/about/reports.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"limit": strconv.Itoa(limit),
	})
	ret := models.ReportListing{}
	ret, err = unMarshalReports(ans, ret)
	return ret.GetChildren(), err
}

func (c *Reddit) getSubredditReportsAfter(sr string, last string, limit int) ([]models.ReportListingChild, error) {
	target := RedditOauth + "/r/" + sr + "/about/reports.json"
	ans, err := c.MiraRequest(http.MethodGet, target, map[string]string{
		"limit":  strconv.Itoa(limit),
		"before": last,
	})
	ret := models.ReportListing{}
	ret, err = unMarshalReports(ans, ret)
	return ret.GetChildren(), err
}
