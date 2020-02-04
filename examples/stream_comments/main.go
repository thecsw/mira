package main

import (
	"github.com/thecsw/mira"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	c, _ := r.Subreddit("all").StreamComments()
	for {
		msg := <-c
		r.Comment(msg.GetId()).Reply("myreply")
	}
}
