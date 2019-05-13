package mira

import (
	"encoding/json"
	"testing"
)

func BenchmarkCreateSubreddit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sub := Subreddit{}
		json.Unmarshal([]byte(orig), &sub)
	}
}
