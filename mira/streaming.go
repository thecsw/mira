package mira

import (
	"time"
)

const (
	CommentListInterval = 8
)

// c is the channel with all unread messages
// stop is the channal to stop the stream. Do stop <- true to stop the loop
func (r *Reddit) StreamCommentReplies() (<-chan ListingDataChildren, chan bool) {
	c := make(chan ListingDataChildren, 25)
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
