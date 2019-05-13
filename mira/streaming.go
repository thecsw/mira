package mira

import (
	"time"
)

func (r *Reddit) StreamCommentReplies() <-chan ListingDataChildren {
	c := make(chan ListingDataChildren, 50)
	go func() {
		for {
			un, _ := r.ListUnreadMessages()
			for _, v := range un.GetMessages() {
				if v.IsCommentReply() {
					c <- v
				}
				r.ReadMessage(v.GetId())
			}
			time.Sleep(5 * time.Second)
		}
	}()
	return c
}
