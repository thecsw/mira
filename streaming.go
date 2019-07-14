package mira

import (
	"time"
)

// c is the channel with all unread messages
// stop is the channel to stop the stream. Do stop <- true to stop the loop
func (r *Reddit) StreamCommentReplies() (<-chan Comment, chan bool) {
	c := make(chan Comment, 25)
	stop := make(chan bool, 1)
	go func() {
		for {
			stop <- false
			un, _ := r.ListUnreadMessages()
			for _, v := range un.GetChildren() {
				// Only process comment replies and
				// mark them as read.
				if v.IsCommentReply() {
					c <- v
					r.ReadMessage(v.GetId())
				}
			}
			time.Sleep(r.Stream.CommentListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop
}

// c is the channel with all comments
// stop is the channel to stop the stream. Do stop <- true to stop the loop
func (r *Reddit) StreamNewComments(sr string) (<-chan Comment, chan bool) {
	c := make(chan Comment, 25)
	stop := make(chan bool, 1)
	anchor, _ := r.GetSubredditComments(sr, "new", "hour", 1)
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			stop <- false
			un, _ := r.GetSubredditCommentsAfter(sr, "new", last, 25)
			for _, v := range un {
				c <- v
			}
			if len(un) > 0 {
				last = un[0].GetId()
			}
			time.Sleep(r.Stream.CommentListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop
}

func (r *Reddit) StreamNewPosts(sr string) (<-chan PostListingChild, chan bool) {
	c := make(chan PostListingChild, 25)
	stop := make(chan bool, 1)
	anchor, _ := r.GetSubredditPosts(sr, "new", "hour", 1)
	last := ""
	if len(anchor.GetChildren()) > 0 {
		last = anchor.GetChildren()[0].GetId()
	}
	go func() {
		for {
			stop <- false
			new, _ := r.GetSubredditPostsAfter(sr, "new", last, r.Stream.PostListSlice)
			s := new.GetChildren()
			if len(s) > 0 {
				last = s[0].GetId()
			}
			for i, _ := range s {
				c <- s[len(s)-i-1]
			}
			time.Sleep(r.Stream.PostListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop
}
