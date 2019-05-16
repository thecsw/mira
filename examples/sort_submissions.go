package main

import (
	"github.com/thecsw/mira"
	"fmt"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	sort := "top"
	var limit int = 25
	subs, _ := r.GetSubredditPosts("subredditname", sort, limit)
	
	for _, v := range subs.GetChildren() {
		fmt.Println("Submission Title: ", v.GetTitle())
	}
}
