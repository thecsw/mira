package mira

import (
	"encoding/json"
	"testing"
)

func BenchmarkCreateMe(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sub := Me{}
		json.Unmarshal([]byte(meExampleJson), &sub)
	}
}
