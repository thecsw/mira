package main

import (
	"fmt"
	"./goraw"
)

func main() {
	cred := goraw.ReadCreds("login.conf")
	// It will handle aute refreshing the tokens
	r := goraw.Init(cred)
	fmt.Println(r)

	fmt.Println(r.Me())
	fmt.Println(r.GetUser("thecsw"))
	post := r.Submit("mypost", "hello, world", "memeinvestor_test")
	comment := post.Comment(r, "My First Comment")
	reply := comment.Reply(r, "My Reply to the First Comment")
	fmt.Println(reply)
}
