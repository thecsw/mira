package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func BenchmarkCreateReportListing(b *testing.B) {
	data, _ := ioutil.ReadFile("./tests/reports.json")
	reportListingExampleJson := string(data)
	for i := 0; i < b.N; i++ {
		sub := ReportListing{}
		json.Unmarshal([]byte(reportListingExampleJson), &sub)
	}
}
