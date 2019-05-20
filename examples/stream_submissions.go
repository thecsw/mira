package main

import (
	"github.com/thecsw/mira"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	c, _ := r.StreamNewPosts("subredditname")
	for {
		post := <-c
		r.Comment(post.GetId(), "I saw your submission!")
	}
}
