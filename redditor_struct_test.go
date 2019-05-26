package mira

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreateRedditor(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/redditor.json")
	redditorExampleJson := string(data)
	for i := 0; i < b.N; i++ {
		sub := Redditor{}
		json.Unmarshal([]byte(redditorExampleJson), &sub)
	}
}
