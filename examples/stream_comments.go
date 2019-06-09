package examples

import (
	"github.com/thecsw/mira"
)

func StreamComments() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	c, _ := r.StreamCommentReplies()
	for {
		msg := <-c
		r.Reply(msg.GetId(), "I got your message!")
	}
}
