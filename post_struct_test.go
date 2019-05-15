package mira

import (
	"encoding/json"
	"testing"
)

func BenchmarkCreatePostListing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sub := PostListing{}
		json.Unmarshal([]byte(postListingExampleJson), &sub)
	}
}
