package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGetReportListingId(t *testing.T) {
	report := ReportListing{}
	data, _ := ioutil.ReadFile("./tests/reports.json")
	json.Unmarshal(data, &report)
	if v := report.GetChildren()[0].GetId(); v != `t1_hw6ng9c` {
		t.Error(
			"For GetId()",
			"expected", `t1_hw6ng9c`,
			"got", v,
		)
	}
}

func TestGetReportListingChildUserReports(t *testing.T) {
	report := ReportListing{}
	data, _ := ioutil.ReadFile("./tests/reports.json")
	json.Unmarshal(data, &report)
	first_post_reports := report.GetChildren()[0].GetUserReports()
	/*
		"user_reports": [
		                        [
		                            "Be Civil",
		                            1,
		                            false,
		                            false
		                        ]
		                    ],
	*/
	if first_post_reports[0].NumOfReports != 1 {
		t.Error(
			"For first_post_reports[0].NumOfReports",
			"expected", 1,
			"got", first_post_reports[0].NumOfReports,
		)
	}
	if first_post_reports[0].SnoozeStatus != false {
		t.Error(
			"For first_post_reports[0].SnoozeStatus",
			"expected", false,
			"got", first_post_reports[0].SnoozeStatus,
		)
	}
	if first_post_reports[0].CanSnooze != false {
		t.Error(
			"For first_post_reports[0].CanSnooze",
			"expected", false,
			"got", first_post_reports[0].CanSnooze,
		)
	}
}
