package mira

import (
	"encoding/json"
	"testing"
)

func BenchmarkCreateCommentListing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sub := CommentListing{}
		json.Unmarshal([]byte(commentListingExampleJson), &sub)
	}
}
