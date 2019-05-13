package main

import (
	"encoding/json"
	"fmt"

	"./goraw"
)

func main() {
	r := goraw.Init(goraw.ReadCredsFromFile("login.conf"))
	post := r.Submit("memeinvestor_test", "mypost", "my text")
	a, _ := json.Marshal(post)
	fmt.Println("POST:", string(a))
	comment := r.Comment(post.GetId(), "My First Comment")
	b, _ := json.Marshal(comment)
	fmt.Println("COMMENT:", string(b))
	reply := r.Reply(comment.GetId(), "My Reply to the First Comment")
	c, _ := json.Marshal(reply)
	fmt.Println("REPLY:", string(c))
}
