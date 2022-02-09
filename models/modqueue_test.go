package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGetModQueueListingId(t *testing.T) {
	modqueue := ModQueueListing{}
	data, _ := ioutil.ReadFile("./tests/modqueue.json")
	json.Unmarshal(data, &modqueue)
	if v := modqueue.GetChildren()[0].GetId(); v != `t1_hw8ecqj` {
		t.Error(
			"For GetId()",
			"expected", `t1_hw8ecqj`,
			"got", v,
		)
	}
}
