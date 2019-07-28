package main

import (
	"github.com/thecsw/mira"
)

func main() {
	// Good practice is to check if the login errors out or not
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	c, _ := r.StreamCommentReplies()
	for {
		msg := <-c
		r.Comment(msg.GetId()).Reply("I got your message!")
	}
}
