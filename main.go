package main

import (
	"fmt"
	"./goraw"
)

func main() {
	r := goraw.Init(goraw.ReadCredsFromFile("login.conf"))
	post := r.Submit("memeinvestor_test", "mypost", "my text")
	comment := r.Comment(post.GetId(), "My First Comment")
	reply := r.Reply(comment.GetId(), "My Reply to the First Comment")
	fmt.Println(reply.GetBody())
}
