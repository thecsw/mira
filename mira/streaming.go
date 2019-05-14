package mira

import (
	"time"
)

const (
	CommentListInterval = 8
	PostListInterval = 5
	PostListSlice = 3
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

func (r *Reddit) StreamNewPosts(sr string) (<-chan Post, chan bool) {
	c := make(chan Post)
	stop := make(chan bool, 1)
	go func(){
		LastTime := time.Now().UTC().Unix()
		for {
			stop <- false
			new, _ := r.GetSubredditPosts(sr, "new", PostListSlice)
			for _, s := range new {
				if s.GetTimeCreated() > float64(LastTime) {
					c <- s.Data
					LastTime = time.Now().UTC().Unix()
				}
			}
			time.Sleep(PostListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop
}
