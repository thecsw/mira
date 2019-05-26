package mira

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreatePostListing(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/postlisting.json")
	postListingExampleJson := string(data)
	for i := 0; i < b.N; i++ {
		sub := PostListing{}
		json.Unmarshal([]byte(postListingExampleJson), &sub)
	}
}
