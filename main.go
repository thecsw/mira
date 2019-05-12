package main

import (
	"fmt"
	"./goraw"
)

func main() {
	cred := goraw.ReadCreds("login.conf")
	// It will handle aute refreshing the tokens
	r := goraw.Init(cred)
//	fmt.Println(r)

//	fmt.Println(r.Me())
//	fmt.Println(r.GetUser("thecsw"))
	r.Submit("memeinvestor_test", "mypost", "my text")
//	comment := post.Comment(r, "My First Comment")
//	reply := r.Reply(comment.GetId(), "My Reply to the First Comment")
//	fmt.Println(reply)
}
