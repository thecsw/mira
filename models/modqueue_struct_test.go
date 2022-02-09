package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreateModQueueListing(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/modqueue.json")
	modQueueListingExampleJson := string(data)
	for i := 0; i < b.N; i++ {
		sub := ModQueueListing{}
		json.Unmarshal([]byte(modQueueListingExampleJson), &sub)
	}
}
