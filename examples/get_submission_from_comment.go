package main

import (
	"fmt"

	"github.com/thecsw/mira"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	ch, stop := r.StreamNewPosts("popular")
	r.Stream.PostListInterval = 2
	i := 0
	for {
		post := <-ch
		i++
		fmt.Println("(", i, ") -->", post.GetTitle())
		comment, _ := r.Comment(post.GetId(), "hello, world")
		reply, _ := r.Reply(comment.GetId(), "you never reply")
		fmt.Println(r.GetSubmissionFromComment(reply.GetId()))
		if i == 1 {
			stop <- true
			break
		}
	}
}
