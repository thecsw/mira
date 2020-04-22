![mira](./mira.png)

[![Go Report Card](https://goreportcard.com/badge/github.com/thecsw/mira)](https://goreportcard.com/report/github.com/thecsw/mira)
[![Build Status](https://travis-ci.org/thecsw/mira.svg?branch=master)](https://travis-ci.org/thecsw/mira)
[![GoDoc](https://godoc.org/github.com/thecsw/mira?status.svg)](https://godoc.org/github.com/thecsw/mira)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

For full documentation, please see the [Godoc page](https://godoc.org/github.com/thecsw/mira)

*mira* is a Reddit Api Wrapper written in beautiful Go.

It is super simple to use the bot as we also provide you with simple but fully extensive 
interfaces. Currently, *mira* is a project that is considered more or less complete. All
main functionality, such as streaming, data manipulation, data request, submissions, links,
etc. are fully implemented. *mira* can be extended to use any Reddit API endpoint. More 
details at the bottom of this page.

## Demo

![Demo](./examples/demo/demo.svg)

Two quick notes: all actions should be done via `Reddit` struct, I thought it would make it
simpler to work with. Secondly, all actions require the objects full `thing_id`, so you have
to use `GetId()` to get that id. Every struct has that method implemented and it will return
a string in the form of `t[1-6]_[a-z0-9]{5}`. Refer to the following table for the classifications
of the structs.

**Type Prefixes**

| Prefix | Type                             |
|--------|----------------------------------|
|   t1   | Comment                          |
|   t2   | Redditor                         |
|   t3   | Submission, PostListing contents |
|   t4   | Message (NOT IMPLEMENTED)        |
|   t5   | Subreddit                        |
|   t6   | Award (NOT IMPLEMENTED)          |

## Config file

The config file structure is very simple:

```
login.conf
----------
CLIENT_ID =
CLIENT_SECRET =
USERNAME =
PASSWORD =
USER_AGENT =
```

``` go
r, err := mira.Init(mira.ReadCredsFromFile("login.conf"))
```

## Environment setup

Mira also works with environmental variables, here is an example from docker-compose

```
    environment:
      - BOT_CLIENT_ID=hunteoahtnhnt432
      - BOT_CLIENT_SECRET=ehoantehont4ht34hnt332
      - BOT_USER_AGENT='u/mytestbot developed by thecsw'
      - BOT_USERNAME=mytestbot
      - BOT_PASSWORD=verygoodpassword
```

And the login will look like this:

``` go
r, err := mira.Init(mira.ReadCredsFromEnv())
```

Or you can always just fill in the values directly.

## Examples

Note: Error checking is omitted for brevity.

### Streaming

Streaming new submissions is very simple! *mira* supports streaming comment replies, 
mentions, new subreddit's/redditor's comments, and new subreddit's/redditor's submissions.

``` go
// r is an instance of *mira.Reddit
r, err := mira.Init(mira.ReadCredsFromFile("login.conf"))

// Start streaming my comment replies
c, err := r.StreamCommentReplies()
for {
	msg := <-c
	r.Comment(msg.GetId()).Reply("I got your message!")
}

// Start streaming my mentions
// Start streaming my comment replies
c, err := r.StreamMentions()
for {
	msg := <-c
	r.Comment(msg.GetId()).Reply("I got your mention of me!")
}

// Start streaming subreddits' submissions
c, err := r.Subreddit("tifu", "wholesomememes").StreamSubmissions()
for {
	post := <-c
	r.Submission(post.GetId()).Save("hello there")
}

// NOTE: Second value is the stop channel. Send a true value
// to the stop channel and the goroutine will return. 
// Basically, `stop <- true`

// Start streaming subreddits' comments
c, err := r.Subreddit("all").StreamComments()
for {
	msg := <-c
	r.Comment(msg.GetId()).Reply("my reply!")
}

// Start streaming redditor's submissions
c, err := r.Redditor("thecsw").StreamSubmissions()
for {
	post := <-c
	r.Submission(post.GetId()).Save("hello there")
}
	
// Start streaming redditor' comments
c, err := r.Redditor("thecsw").StreamComments()
for {
	msg := <-c
	r.Comment(msg.GetId()).Reply("my reply!")
}
```

### Submitting, Commenting, Replying, and Editing

It is very easy to post a submission, comment on it, reply to a message, or
edit a comment.

``` go
package main

import (
	"fmt"

	"github.com/thecsw/mira"
)

// Error checking is omitted for brevity
func main() {
	r, err := mira.Init(mira.ReadCredsFromFile("login.conf"))

	// Make a submission
	post, err := r.Subreddit("mysubreddit").Submit("mytitle", "mytext")

	// Comment on our new submission
	comment, err := r.Submission(post.GetId()).Save("mycomment")

	// Reply to our own comment
	reply, err := r.Comment(comment.GetId()).Reply("myreply")

	// Delete the reply
	r.Comment(reply.GetId()).Delete()

	// Edit the first comment
	newComment, err := r.Comment(comment.GetId()).Edit("myedit")

	// Show the comment's body
	fmt.Println(newComment.GetBody())
}
```

### Composing a message

We can also send a message to another user!

``` go
package main

import (
	"github.com/thecsw/mira"
)

func main() {
	r, err := mira.Init(mira.ReadCredsFromFile("login.conf"))

	r.Redditor("myuser").Compose("mytitle", "mytext")
}
```

### Going through hot, new, top, rising, controversial, and random

You can also traverse through a number of submissions using
one of our methods.

``` go
package main

import (
	"fmt"

	"github.com/thecsw/mira"
)

func main() {
	r, err := mira.Init(mira.ReadCredsFromFile("login.conf"))
	sort := "top"
	var limit int = 25
	duration := "all"
	subs, err := r.Subreddit("all").Submissions(sort, duration, limit)
	for _, v := range subs {
		fmt.Println("Submission Title: ", v.GetTitle())
	}
}
```

### Getting reddit info

You can extract info from any reddit ID using mira. The returned value is an 
instance of mira.MiraInterface.

``` go
package main

import (
	"fmt"

	"github.com/thecsw/mira"
)

func main() {
	r, err := mira.Init(mira.ReadCredsFromFile("login.conf"))
	me, err := r.Me().Info()
	comment, err := r.Comment("t1_...").Info()
	redditor, err := r.Redditor.Info("t2_...")
	submission, err := r.Submission("t3_...").Info()
	subreddit, err := r.Subreddit("t5_...").Info()
}
```

Here is the interface:

``` go
type MiraInterface interface {
	GetId() string
	GetParentId() string
	GetTitle() string
	GetBody() string
	GetAuthor() string
	GetName() string
	GetKarma() float64
	GetUps() float64
	GetDowns() float64
	GetSubreddit() string
	GetCreated() float64
	GetFlair() string
	GetUrl() string
	IsRoot() bool
}
```

## Mira Caller

Surely, Reddit API is always developing and I can't implement all endpoints. It will be a bit of a bloat.
Instead, you have accessto *Reddit.MiraRequest method that will let you to do any custom reddit api calls!

Here is the signature:

``` go
func (c *Reddit) MiraRequest(method string, target string, payload map[string]string) ([]byte, error) {...}
```

It is pretty straight-forward. The return is a slice of bytes. Parse it yourself.

Here is an example of how Reddit.Reply() uses MiraRequest:

NOTE: `checkType(...)` is a quick method to pop a value from the
queue and make sure it's a valid value and type. For example,

``` go
r.Comment("COMM1").Submission("SUBM1").Redditor("USER1")
```

will add elements to its internal queue, so that the layout is:

```
Enqueue->
              redditor  submission  comment               // type
	|BACK| -> |USER1| -> |SUBM1| -> |COMM1| -> |FRONT|    // value
                                                  Dequeue->
```

So that when you run `r.checkType("comment")`, it will dequeue `COMM1`
and return triplet `"COMM1", "comment", nil`.

If you run `r.checkType("redditor")` (will fail because subm is at the end),
you will get `"", "", "errors.New("the passed type...")`

Here is an example of how you check that the last element to dequeue is
a type that you're expecting:

``` go
func (c *Reddit) Reply(text string) (models.CommentWrap, error) {
	ret := &models.CommentWrap{}
	// Second return is type, which is "comment"
	name, _, err := c.checkType("comment")
	if err != nil {
		return *ret, err
	}
	target := RedditOauth + "/api/comment"
	ans, err := c.MiraRequest("POST", target, map[string]string{
		"text":     text,
		"thing_id": name,
		"api_type": "json",
	})
	json.Unmarshal(ans, ret)
	return *ret, err
}
```
