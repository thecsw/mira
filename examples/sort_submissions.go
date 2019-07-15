package examples

import (
	"fmt"

	"github.com/thecsw/mira"
)

func SortSubmissions() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))
	sort := "top"
	var limit int = 25
	duration := "all"
	subs, _ := r.GetSubredditPosts("subredditname", sort, duration, limit)

	for _, v := range subs {
		fmt.Println("Submission Title: ", v.GetTitle())
	}
}
