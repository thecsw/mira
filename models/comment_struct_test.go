package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreateComment(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/comment.json")
	commentExampleJson := string(data)
	for i := 0; i < b.N; i++ {
		sub := CommentWrap{}
		json.Unmarshal([]byte(commentExampleJson), &sub)
	}
}
