package mira

import (
	"time"
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
	last := ""
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
