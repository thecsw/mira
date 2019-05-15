package main

import (
	"./mira"
	"fmt"
//	"encoding/json"
//	"time"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))

	post, _ := r.Submit("memeinvestor_test", "mypost", "my text")
	comment, _ := r.Comment(post.GetId(), "My First Comment")
	reply, _ := r.Reply(comment.GetId(), "My Reply to the First Comment")
	r.DeleteComment(reply.GetId())
	fmt.Println(comment.GetBody())
	new_comment, _ := r.EditComment(comment.GetId(), "I Edited This!!")
	fmt.Println(new_comment.GetBody())

	// arr, _ := r.GetSubredditPosts("memeeconomy", "top", 10)
	// fmt.Println(arr)

	c, _ := r.StreamNewPosts("memeeconomy")
	for {
		data := <- c
		fmt.Println("Title: ", data.GetTitle())
	}
}
