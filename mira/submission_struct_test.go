package mira

import (
	"encoding/json"
	"testing"
)

func BenchmarkCreateSubmission(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sub := Submission{}
		json.Unmarshal([]byte(submissionExampleJson), &sub)
	}
}
