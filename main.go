package main

import (
	"./goraw"
	"fmt"
	"encoding/json"
)

func main() {
	r := goraw.Init(goraw.ReadCredsFromFile("login.conf"))

	// Testing if the token refresh actually works
	// The default expiration is 60 minutes, I will run
	// this for 100, so if it works, it will continue
	// posting. Otherwise, it will throw an error
	post, _ := r.Submit("memeinvestor_test", "mypost", "my text")
	comment, _ := r.Comment(post.GetId(), "My First Comment")
	reply, _ := r.Reply(comment.GetId(), "My Reply to the First Comment")
	r.DeleteComment(reply.GetId())
	fmt.Println(comment.GetBody())
	new_comment, _ := r.EditComment(comment.GetId(), "I Edited This!!")
	fmt.Println(new_comment.GetBody())

	// thecsw, _ := r.GetUser("thecsw")
	// out, _ := json.Marshal(thecsw)
	// fmt.Println(string(out))
	// r.Compose("Thecsw", "hello, world", "Can you see this?")
	//	ll, _ := r.ReadAllMessages()
	//	fmt.Println(string(ll))	
	un, _ := r.ListUnreadMessages()
	for i, v := range un.GetMessages() {
		data, _ := json.Marshal(v)
		fmt.Println(string(data))
		fmt.Println(i, v, v.IsComment())
		r.ReadMessage(v.GetId())
	}
}
