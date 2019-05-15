package mira

import (
	"encoding/json"
	"testing"
)

func BenchmarkCreateComment(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sub := Comment{}
		json.Unmarshal([]byte(commentExampleJson), &sub)
	}
}
