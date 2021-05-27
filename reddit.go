package mira

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/thecsw/mira/models"
)

const (
	submissionType = "s"
	subredditType  = "b"
	commentType    = "c"
	redditorType   = "r"
	meType         = "m"
)

// MiraRequest Reddit API is always developing and I can't implement all endpoints;
// It will be a bit of a bloat; Instead, you have accessto *Reddit.MiraRequest
// method that will let you to do any custom reddit api calls!
//
// Here is the signature:
//
//   func (c *Reddit) MiraRequest(method string, target string, payload map[string]string) ([]byte, error) {...}
//
// It is pretty straight-forward, the return is a slice of bytes; Parse it yourself.
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

// Me pushes a new Redditor value.
func (c *Reddit) Me() *Reddit {
	return c.addQueue(c.Creds.Username, meType)
}

// Subreddit pushes a new subreddit value to the internal queue.
func (c *Reddit) Subreddit(name ...string) *Reddit {
	return c.addQueue(strings.Join(name, "+"), subredditType)
}

// Submission pushes a new submission value to the internal queue.
func (c *Reddit) Submission(name string) *Reddit {
	return c.addQueue(name, submissionType)
}

// Comment pushes a new comment value to the internal queue.
func (c *Reddit) Comment(name string) *Reddit {
	return c.addQueue(name, commentType)
}

// Redditor pushes a new redditor value to the internal queue.
func (c *Reddit) Redditor(name string) *Reddit {
	return c.addQueue(name, redditorType)
}

// Submissions returns submissions from a subreddit up to a specified limit sorted by the given parameters
//
// Sorting options: "hot", "new", "top", "rising", "controversial", "random"
//
// Duration options: "hour", "day", "week", "year", "all"
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
// Last is the anchor of a submission id
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

// Info returns MiraInterface of last pushed object.
func (c *Reddit) Info() (MiraInterface, error) {
	name, ttype := c.getQueue()
	switch ttype {
	case meType:
		return c.getMe()
	case submissionType:
		return c.getSubmission(name)
	case commentType:
		return c.getComment(name)
	case subredditType:
		return c.getSubreddit(name)
	case redditorType:
		return c.getUser(name)
	default:
		return nil, fmt.Errorf("returning type is not defined")
	}
}

func (c *Reddit) getMe() (models.Me, error) {
	target := RedditOauth + "/api/v1/me"
	ret := &models.Me{}
	ans, err := c.MiraRequest("GET", target, nil)
	if err != nil {
		return *ret, err
	}
	json.Unmarshal(ans, ret)
	return *ret, nil
}

func (c *Reddit) getSubmission(id string) (models.PostListingChild, error) {
	target := RedditOauth + "/api/info.json"
	ans, err := c.MiraRequest("GET", target, map[string]string{
		"id": id,
	})
	ret := &models.PostListing{}
	json.Unmarshal(ans, ret)
	if len(ret.GetChildren()) < 1 {
		return models.PostListingChild{}, fmt.Errorf("id not found")
	}
	return ret.GetChildren()[0], err
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

func (c *Reddit) getUser(name string) (models.Redditor, error) {
	target := RedditOauth + "/user/" + name + "/about"
	ans, err := c.MiraRequest(http.MethodGet, target, nil)
	ret := &models.Redditor{}
	json.Unmarshal(ans, ret)
	return *ret, err
}

func (c *Reddit) getSubreddit(name string) (models.Subreddit, error) {
	target := RedditOauth + "/r/" + name + "/about"
	ans, err := c.MiraRequest(http.MethodGet, target, nil)
	ret := &models.Subreddit{}
	json.Unmarshal(ans, ret)
	return *ret, err
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
		"api_type": "json",
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

func (c *Reddit) checkType(rtype ...string) (string, string, error) {
	name, ttype := c.getQueue()
	if name == "" {
		return "", "", fmt.Errorf("identifier is empty")
	}
	if !findElem(ttype, rtype) {
		return "", "", fmt.Errorf(
			"the passed type is not a valid type for this call | expected: %s",
			strings.Join(rtype, ", "))
	}
	return name, ttype, nil
}

func (c *Reddit) addQueue(name string, ttype string) *Reddit {
	c.Chain <- &ChainVals{Name: name, Type: ttype}
	return c
}

func (c *Reddit) getQueue() (string, string) {
	if len(c.Chain) < 1 {
		return "", ""
	}
	temp := <-c.Chain
	return temp.Name, temp.Type
}

func findElem(elem string, arr []string) bool {
	for _, v := range arr {
		if elem == v {
			return true
		}
	}
	return false
}

// RedditErr is a struct to store reddit error messages.
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
