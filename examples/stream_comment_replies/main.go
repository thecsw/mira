package main

import (
	"github.com/thecsw/mira/v3"
)

func main() {
	// Good practice is to check if the login errors out or not
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	c := r.StreamCommentReplies()
	for {
		msg := <-c
		r.Comment(msg.GetId()).Reply("I got your message!")
	}
}
