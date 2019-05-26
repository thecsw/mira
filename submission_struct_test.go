package mira

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreateSubmission(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/submission.json")
	submissionExampleJson := string(data)
	for i := 0; i < b.N; i++ {
		sub := Submission{}
		json.Unmarshal([]byte(submissionExampleJson), &sub)
	}
}
