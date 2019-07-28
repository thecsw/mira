package main

import (
	"fmt"

	"github.com/thecsw/mira"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	sort := "top"
	var limit int = 25
	duration := "all"
	subs, _ := r.Subreddit("all").Submissions(sort, duration, limit)
	for _, v := range subs {
		fmt.Println("Submission Title: ", v.GetTitle())
	}
}
