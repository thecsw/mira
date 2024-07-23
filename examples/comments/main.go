package main

import (
	"fmt"

	"github.com/thecsw/mira/v3"
)

// Errors are omitted for brevity
func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))

	// Make a submission
	post, _ := r.Subreddit("mysubreddit").Submit("mytitle", "mytext")

	// Comment on our new submission
	comment, _ := r.Submission(post.GetId()).Save("mycomment")

	// Reply to our own comment
	reply, _ := r.Comment(comment.GetId()).Reply("myreply")

	// Delete the reply
	r.Comment(reply.GetId()).Delete()

	// Edit the first comment
	newComment, _ := r.Comment(comment.GetId()).Edit("myedit")

	// Show the comment's body
	fmt.Println(newComment.GetBody())
}
