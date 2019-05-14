package mira

import (
	"testing"
	"encoding/json"
)

func BenchmarkCreateRedditor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sub := Redditor{}
		json.Unmarshal([]byte(redditorExampleJson), &sub)
	}
}
