# MIRA or Meme Investor Reddit Api

[![Go Report Card](https://goreportcard.com/badge/github.com/thecsw/mira)](https://goreportcard.com/report/github.com/thecsw/mira)
[![Build Status](https://travis-ci.org/thecsw/mira.svg?branch=master)](https://travis-ci.org/thecsw/mira)
[![GoDoc](https://godoc.org/github.com/thecsw/mira?status.svg)](https://godoc.org/github.com/thecsw/mira)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)

For full documentation, please see the [Godoc page](https://godoc.org/github.com/thecsw/mira)

MIRA is a Reddit Api Wrapper written in beautiful Go. This is a subproject
of a bigger project MemeInvestor_bot. 

It is super simple to use the bot as we also provide you
with simple but fully extensive structs. We utilize the 
best features of Go, such as closures, channels, goroutines, garbage collection, etc.

Currently, `mira` is a project that just began its life. We still have many new Reddit
endpoints to cover. We have the basic functionality implemented, like streaming comment
replies, new submissions. Comment, post, edit, reply, and delete options for our
instances.

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
CLIENT_ID =
CLIENT_SECRET =
USERNAME =
PASSWORD =
USER_AGENT =
```

## Examples

### Streaming comment replies

Below is an example on how to make a simple bot that 
listens to a stream of comment replies and replies.

``` go
package main

import (
	"github.com/thecsw/mira"
	"fmt"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	c, _ := r.StreamCommentReplies()
	for {
		msg := <- c
		r.Reply(msg.GetId(), "I got your message!")
	}
}
```

### Streaming new submissions

Streaming new submissions is very simple too. You can do it the same way
as streaming comment replies.

``` go
package main

import (
	"github.com/thecsw/mira"
	"fmt"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	c, _ := r.StreamNewPosts("subredditname")
	for {
		post := <- c
		r.Comment(post.GetId(), "I saw your submission!")
	}
}
```

### Submitting, Commenting, Replying, and Editing

It is very easy to post a submission, comment on it, reply to a message, or
edit a comment.

``` go
package main

import (
	"github.com/thecsw/mira"
	"fmt"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	// Make a submission
	post, _ := r.Submit("memeinvestor_test", "mypost", "my text")
	// Comment on our new submission
	comment, _ := r.Comment(post.GetId(), "My First Comment")
	// Reply to our own comment
	reply, _ := r.Reply(comment.GetId(), "My Reply to the First Comment")
	// Delete the reply
	r.DeleteComment(reply.GetId())
	// Edit the first comment
	new_comment, _ := r.EditComment(comment.GetId(), "I Edited This!!")
	// Show the comment's body
	fmt.Println(new_comment.GetBody())
}
```

### Composing a message

We can also send a message to another user!

``` go
package main

import (
	"github.com/thecsw/mira"
	"fmt"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))

	r.Compose("thecsw", "my subject", "hello, world")
	// or
	user, _ := r.GetUser("thecsw")
	r.Compose(user.GetName(), "my subject", "hello, world")
}
```

### Going through hot, new, top, rising, controversial, and random

You can also traverse through a number of submissions using
one of our methods.

``` go
package main

import (
	"github.com/thecsw/mira"
	"fmt"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	sort := "top"
	var limit int = 25
	duration := "all"
	subs, _ := r.GetSubredditPosts("subredditname", sort, duration, limit)
	
	for _, v := range subs.GetChildren() {
		fmt.Println("Submission Title: ", v.GetTitle())
	}
}
```
