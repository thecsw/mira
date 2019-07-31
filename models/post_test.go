package models

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGetPostListingId(t *testing.T) {
	post := PostListing{}
	data, _ := ioutil.ReadFile("./tests/postlisting.json")
	json.Unmarshal(data, &post)
	if v := post.GetChildren()[0].GetId(); v != `t3_bev1v7` {
		t.Error(
			"For GetId()",
			"expected", `t3_bev1v7`,
			"got", v,
		)
	}
}

func TestGetSubreddit(t *testing.T) {
	post := PostListing{}
	data, _ := ioutil.ReadFile("./tests/postlisting.json")
	json.Unmarshal(data, &post)
	if v := post.GetChildren()[0].GetSubreddit(); v != `MemeEconomy` {
		t.Error(
			"For GetSubreddit()",
			"expected", `MemeEconomy`,
			"got", v,
		)
	}
}

func TestGetTitle(t *testing.T) {
	post := PostListing{}
	data, _ := ioutil.ReadFile("./tests/postlisting.json")
	json.Unmarshal(data, &post)
	if v := post.GetChildren()[1].GetTitle(); v != `Slow it down a bit, and invest here for THICC profits` {
		t.Error(
			"For GetTitle()",
			"expected", `Slow it down a bit, and invest here for THICC profits`,
			"got", v,
		)
	}
}
