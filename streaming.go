package mira

import (
	"time"
	//	"fmt"
)

var (
	// Comment Replies are more frequent, every 8 seconds should be fine
	CommentListInterval = 8
	// Submissions are a bit more rare. To save the API limit, every 15 seconds should be enough
	PostListInterval = 10
	// Just to keep everything, size of 8 rounds up good
	PostListSlice = 8
)

// c is the channel with all unread messages
// stop is the channel to stop the stream. Do stop <- true to stop the loop
func (r *Reddit) StreamCommentReplies() (<-chan CommentListingDataChildren, chan bool) {
	c := make(chan CommentListingDataChildren, 25)
	stop := make(chan bool, 1)
	go func() {
		for {
			stop <- false
			un, _ := r.ListUnreadMessages()
			for _, v := range un.GetMessages() {
				// Only process comment replies and
				// mark them as read.
				if v.IsCommentReply() {
					c <- v
					r.ReadMessage(v.GetId())
				}
			}
			time.Sleep(CommentListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop
}

func (r *Reddit) StreamNewPosts(sr string) (<-chan PostListingChild, chan bool) {
	c := make(chan PostListingChild)
	stop := make(chan bool, 1)
	last := ""
	go func() {
		for {
			stop <- false
			new, _ := r.GetSubredditPostsAfter(sr, "new", last, PostListSlice)
			s := new.GetChildren()
			if len(s) > 0 {
				last = s[0].GetId()
			}
			for i := len(new.GetChildren()) - 1; i >= 0; i-- {
				c <- s[i]
			}
			time.Sleep(PostListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop
}
