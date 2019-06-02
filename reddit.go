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

func (c *Reddit) Me() (Me, error) {
	target := RedditOauth + "/api/v1/me"
	user := Me{}
	r, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return user, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return user, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &user)
	return user, nil
}

func (c *Reddit) GetComment(id string) (CommentListingDataChildren, error) {
	target := RedditOauth + "/api/info.json?id=" + id
	list := CommentListing{}
	temp := CommentListingDataChildren{}
	r, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return temp, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return temp, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &list)
	return list.GetChildren()[0], nil
}

func (c *Reddit) GetSubmission(id string) (PostListing, error) {
	target := RedditOauth + "/api/info.json?id=" + id
	list := PostListing{}
	r, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return list, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return list, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &list)
	return list, nil
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
		return "", errors.New("The passed ID is not a comment.")
	}
	target := RedditOauth + "/api/info.json?id="
	temp := CommentListing{}
	tries := 0
	for string(current[1]) != "3" {
		r, err := http.NewRequest("GET", target+current, nil)
		if err != nil {
			return "", err
		}
		r.Header.Set("User-Agent", c.Creds.UserAgent)
		r.Header.Set("Authorization", "bearer "+c.Token)
		client := &http.Client{}
		response, err := client.Do(r)
		if err != nil {
			return "", err
		}
		defer response.Body.Close()
		buf := new(bytes.Buffer)
		buf.ReadFrom(response.Body)
		json.Unmarshal(buf.Bytes(), &temp)
		current = temp.Data.Children[0].GetParentId()
		tries++
		if tries > c.Values.GetSubmissionFromCommentTries {
			return "", errors.New(fmt.Sprintf("Exceeded the maximum number of iterations: %v", c.Values.GetSubmissionFromCommentTries))
		}
	}
	return current, nil
}

func (c *Reddit) GetUser(name string) (Redditor, error) {
	target := RedditOauth + "/user/" + name + "/about"
	user := Redditor{}
	r, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return user, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return user, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &user)
	return user, nil
}

func (c *Reddit) GetSubreddit(name string) (Subreddit, error) {
	target := RedditOauth + "/r/" + name + "/about"
	sub := Subreddit{}
	r, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return sub, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return sub, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &sub)
	return sub, nil
}

// Get submisssions from a subreddit up to a specified limit sorted by the given parameter
//
// Sorting options: "hot", "new", "top", "rising", "controversial", "random"
//
// Time options: "all", "year", "month", "week", "day", "hour"
//
// Limit is any numerical value, so 0 <= limit <= 100
func (c *Reddit) GetSubredditPosts(sr string, sort string, tdur string, limit int) (PostListing, error) {
	target := RedditOauth + "/r/" + sr + "/" + sort + ".json" + "?limit=" + strconv.Itoa(limit) + "&t=" + tdur
	listing := PostListing{}
	r, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return listing, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return listing, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &listing)
	return listing, nil
}

func (c *Reddit) GetSubmissionComments(sr string, post_id string, sort string, limit int) ([]CommentListingDataChildren, error) {
	if string(post_id[1]) != "3" {
		return nil, errors.New("The passed ID36 is not a submission.")
	}
	target := RedditOauth + "/r/" + sr + "/comments/" + post_id[3:] + "?sort=" + sort + "&limit=" + strconv.Itoa(limit) + "&showmore=true"
	fmt.Println(target)
	listing := make([]CommentListing, 2)
	//	more:= make([]MoreListing, 10)
	r, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &listing)
	//	json.Unmarshal(buf.Bytes(), &more)
	//	fmt.Println(more)
	return listing[1].GetChildren()[:len(listing[1].GetChildren())-1], nil
}

// Get submisssions from a subreddit up to a specified limit sorted by the given parameters
// and with specified anchor
//
// Sorting options: "hot", "new", "top", "rising", "controversial", "random"
//
// Limit is any numerical value, so 0 <= limit <= 100
//
// Anchor options are submissions full thing, for example: t3_bqqwm3
func (c *Reddit) GetSubredditPostsAfter(sr string, sort string, last string, limit int) (PostListing, error) {
	target := RedditOauth + "/r/" + sr + "/" + sort + ".json" + "?limit=" + strconv.Itoa(limit) + "&before=" + last
	listing := PostListing{}
	r, err := http.NewRequest("GET", target, nil)
	if err != nil {
		return listing, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return listing, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &listing)
	return listing, nil
}

func (c *Reddit) Submit(sr string, title string, text string) (Submission, error) {
	target := RedditOauth + "/api/submit"
	post := Submission{}
	form := url.Values{}
	form.Add("title", title)
	form.Add("sr", sr)
	form.Add("text", text)
	form.Add("kind", "self")
	form.Add("resubmit", "true")
	form.Add("api_type", "json")
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return post, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return post, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &post)
	return post, nil
}

func (c *Reddit) Reply(comment_id string, text string) (Comment, error) {
	target := RedditOauth + "/api/comment"
	comment := Comment{}
	form := url.Values{}
	form.Add("text", text)
	form.Add("thing_id", comment_id)
	form.Add("api_type", "json")
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return comment, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return comment, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &comment)
	return comment, nil
}

func (c *Reddit) Comment(submission_id, text string) (Comment, error) {
	target := RedditOauth + "/api/comment"
	comment := Comment{}
	form := url.Values{}
	form.Add("text", text)
	form.Add("thing_id", submission_id)
	form.Add("api_type", "json")
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return comment, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return comment, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &comment)
	return comment, nil
}

func (c *Reddit) DeleteComment(comment_id string) error {
	target := RedditOauth + "/api/del"
	form := url.Values{}
	form.Add("id", comment_id)
	form.Add("api_type", "json")
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func (c *Reddit) Approve(comment_id string) error {
	target := RedditOauth + "/api/approve"
	form := url.Values{}
	form.Add("id", comment_id)
	form.Add("api_type", "json")
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func (c *Reddit) Distinguish(comment_id string, how string, sticky bool) error {
	st := "false"
	if sticky {
		st = "true"
	}
	target := RedditOauth + "/api/distinguish"
	form := url.Values{}
	form.Add("id", comment_id)
	form.Add("how", how)
	form.Add("sticky", st)
	form.Add("api_type", "json")
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func (c *Reddit) EditComment(comment_id, text string) (Comment, error) {
	target := RedditOauth + "/api/editusertext"
	comment := Comment{}
	form := url.Values{}
	form.Add("text", text)
	form.Add("thing_id", comment_id)
	form.Add("api_type", "json")
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return comment, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return comment, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &comment)
	return comment, nil
}

func (c *Reddit) Compose(to, subject, text string) error {
	target := RedditOauth + "/api/compose"
	form := url.Values{}
	form.Add("subject", subject)
	form.Add("text", text)
	form.Add("to", to)
	form.Add("api_type", "json")
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func (c *Reddit) ReadMessage(message_id string) error {
	target := RedditOauth + "/api/read_message"
	form := url.Values{}
	form.Add("id", message_id)
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func (c *Reddit) ReadAllMessages() error {
	target := RedditOauth + "/api/read_all_messages"
	r, err := http.NewRequest("POST", target, nil)
	if err != nil {
		return err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

func (c *Reddit) ListUnreadMessages() (CommentListing, error) {
	target := RedditOauth + "/message/unread"
	list := CommentListing{}
	form := url.Values{}
	form.Add("mark", "true")
	r, err := http.NewRequest("GET", target, strings.NewReader(form.Encode()))
	if err != nil {
		return list, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return list, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	json.Unmarshal(buf.Bytes(), &list)
	return list, nil
}

func (c *Reddit) SubredditUpdateSidebar(sr, text string) ([]byte, error) {
	target := RedditOauth + "/api/site_admin"
	form := url.Values{}
	form.Add("sr", sr)
	form.Add("name", "None")
	form.Add("description", text)
	form.Add("title", sr)
	form.Add("wikimode", "anyone")
	form.Add("link_type", "any")
	form.Add("type", "public")
	form.Add("api_type", "json")
	r, err := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	return buf.Bytes(), nil
}
