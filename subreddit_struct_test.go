package mira

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreateSubreddit(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	subredditExampleJson := string(data)
	for i := 0; i < b.N; i++ {
		sub := Subreddit{}
		json.Unmarshal([]byte(subredditExampleJson), &sub)
	}
}
