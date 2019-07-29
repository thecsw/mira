package main

import (
	"github.com/thecsw/mira"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	c, _, _ := r.Subreddit("all").StreamSubmissions()
	for {
		post := <-c
		r.Submission(post.GetId()).Save("hello there")
	}
}
