package main

import (
	"github.com/thecsw/mira"
	"fmt"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	// Make a submission
	post, _ := r.Submit("memeinvestor_test", "mypost", "my text")
	// Comment on our new submission
	comment, _ := r.Comment(post.GetId(), "My First Comment")
	// Reply to our own comment
	reply, _ := r.Reply(comment.GetId(), "My Reply to the First Comment")
	// Delete the reply
	r.DeleteComment(reply.GetId())
	// Edit the first comment
	new_comment, _ := r.EditComment(comment.GetId(), "I Edited This!!")
	// Show the comment's body
	fmt.Println(new_comment.GetBody())
}
