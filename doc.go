// It is super simple to use the bot as we also provide you
// with simple but fully extensive structs. We utilize the
// best features of Go, such as closures, channels, goroutines, garbage collection, etc.
//
// Currently, `mira` is a project that just began its life. We still have many new Reddit
// endpoints to cover. We have the basic functionality implemented, like streaming comment
// replies, new submissions. Comment, post, edit, reply, and delete options for our
// instances.
//
// Two quick notes: all actions should be done via `Reddit` struct, I thought it would make it
// simpler to work with. Secondly, all actions require the objects full `thing_id`, so you have
// to use `GetId()` to get that id. Every struct has that method implemented and it will return
// a string in the form of `t[1-6]_[a-z0-9]{5}`. Refer to the following table for the classifications
// of the structs.
//
// **Type Prefixes**
//
// | Prefix | Type                             |
// |--------|----------------------------------|
// |   t1   | Comment                          |
// |   t2   | Redditor                         |
// |   t3   | Submission, PostListing contents |
// |   t4   | Message (NOT IMPLEMENTED)        |
// |   t5   | Subreddit                        |
// |   t6   | Award (NOT IMPLEMENTED)          |
//
// ## Config file
//
// The config file structure is very simple:
//
// ```
// CLIENT_ID =
// CLIENT_SECRET =
// USERNAME =
// PASSWORD =
// USER_AGENT =
// ```
//
// ## Examples
//
// Note: Error checking is omitted for brevity.
//
// ### Streaming comment replies
//
// Below is an example on how to make a simple bot that
// listens to a stream of comment replies and replies.
//
// ``` go
// package main
//
// import (
// 	"github.com/thecsw/mira"
// )
//
// func main() {
// 	// Good practice is to check if the login errors out or not
// 	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
// 	c, _ := r.StreamCommentReplies()
// 	for {
// 		msg := <-c
// 		r.Comment(msg.GetId()).Reply("I got your message!")
// 	}
// }
// ```
//
// ### Streaming new submissions
//
// Streaming new submissions is very simple too. You can do it the same way
// as streaming comment replies.
//
// ``` go
// package main
//
// import (
// 	"github.com/thecsw/mira"
// )
//
// func main() {
// 	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
// 	c, _, _ := r.Subreddit("all").StreamSubmissions()
// 	for {
// 		post := <-c
// 		r.Submission(post.GetId()).Save("hello there")
// 	}
// }
// ```
//
// ### Submitting, Commenting, Replying, and Editing
//
// It is very easy to post a submission, comment on it, reply to a message, or
// edit a comment.
//
// ``` go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/thecsw/mira"
// )
//
// // Errors are omitted for brevity
// func main() {
// 	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
//
// 	// Make a submission
// 	post, _ := r.Subreddit("mysubreddit").Submit("mytitle", "mytext")
//
// 	// Comment on our new submission
// 	comment, _ := r.Submission(post.GetId()).Save("mycomment")
//
// 	// Reply to our own comment
// 	reply, _ := r.Comment(comment.GetId()).Reply("myreply")
//
// 	// Delete the reply
// 	r.Comment(reply.GetId()).Delete()
//
// 	// Edit the first comment
// 	newComment, _ := r.Comment(comment.GetId()).Edit("myedit")
//
// 	// Show the comment's body
// 	fmt.Println(newComment.GetBody())
// }
// ```
//
// ### Composing a message
//
// We can also send a message to another user!
//
// ``` go
// package main
//
// import (
// 	"github.com/thecsw/mira"
// )
//
// func main() {
// 	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
//
// 	r.Redditor("myuser").Compose("mytitle", "mytext")
// }
// ```
//
// ### Going through hot, new, top, rising, controversial, and random
//
// You can also traverse through a number of submissions using
// one of our methods.
//
// ``` go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/thecsw/mira"
// )
//
// func main() {
// 	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
// 	sort := "top"
// 	var limit int = 25
// 	duration := "all"
// 	subs, _ := r.Subreddit("all").Submissions(sort, duration, limit)
// 	for _, v := range subs {
// 		fmt.Println("Submission Title: ", v.GetTitle())
// 	}
// }
// ```
package mira

import (
	// models package
	_ "github.com/thecsw/mira/models"
)
