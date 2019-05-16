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
