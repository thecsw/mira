package mira

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestGetId(t *testing.T) {
	sub := Subreddit{}
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	json.Unmarshal(data, &sub)
	if v := sub.GetId(); v != `t5_m0je4` {
		t.Error(
			"For GetId()",
			"expected", `t5_m0je4`,
			"got", v,
		)
	}
}

func TestGetName(t *testing.T) {
	sub := Subreddit{}
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	json.Unmarshal(data, &sub)
	if v := sub.GetName(); v != `MemeInvestor_bot` {
		t.Error(
			"For GetName()",
			"expected", `MemeInvestor_bot`,
			"got", v,
		)
	}
}

func TestGetDisplayName(t *testing.T) {
	sub := Subreddit{}
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	json.Unmarshal(data, &sub)
	if v := sub.GetDisplayName(); v != `MemeInvestor_bot` {
		t.Error(
			"For GetDisplayName()",
			"expected", `MemeInvestor_bot`,
			"got", v,
		)
	}
}

func TestGetUrl(t *testing.T) {
	sub := Subreddit{}
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	json.Unmarshal(data, &sub)
	if v := sub.GetUrl(); v != `/r/MemeInvestor_bot/` {
		t.Error(
			"For GetUrl()",
			"expected", "/r/MemeInvestor_bot/",
			"got", v,
		)
	}
}

func TestIsOver18(t *testing.T) {
	sub := Subreddit{}
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	json.Unmarshal(data, &sub)
	if v := sub.IsOver18(); v != false {
		t.Error(
			"For IsOver18()",
			"expected", false,
			"got", v,
		)
	}
}

func TestGetPublicDescription(t *testing.T) {
	sub := Subreddit{}
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	json.Unmarshal(data, &sub)
	if v := sub.GetPublicDescription(); v != "This subreddit is for questions, reports, or suggestions regarding /u/MemeInvestor_Bot. \n\nFor quick information see https://memes.market" {
		t.Error(
			"For GetPublicDescription()",
			"expected", "This subreddit is for questions, reports, or suggestions regarding /u/MemeInvestor_Bot. \n\nFor quick information see https://memes.market",
			"got", v,
		)
	}
}

func TestGetDescription(t *testing.T) {
	sub := Subreddit{}
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	json.Unmarshal(data, &sub)
	if v := sub.GetDescription(); v != "######This is the official subreddit for the bot, /u/MemeInvestor_bot\n\n***\n\nHere you're encouraged to **report bugs, ask questions, and submit suggestions** regarding the bot. This subreddit is frequently viewed by the developers and, whether or not you receive a reply, it's very likely that your submission has been viewed and noted by someone on the team.\n\n***\n\n#####Rules:\n1. This is a no-meme subreddit. Only serious suggestions, reports, or questions allowed.\n\n2.  All content must be regarding the bot. Keep it on-topic please.\n\n3. Be respectful. We're all nice people here.\n\n***\n\n&amp;nbsp;\n\n####^(Please don't send a message before first submitting your post on the subreddit.)\n\n######**[Message us anyway.](https://www.reddit.com/message/compose?to=%2Fr%2FMemeInvestor_Bot)**" {
		t.Error(
			"For GetDescription()",
			"expected", "######This is the official subreddit for the bot, /u/MemeInvestor_bot\n\n***\n\nHere you're encouraged to **report bugs, ask questions, and submit suggestions** regarding the bot. This subreddit is frequently viewed by the developers and, whether or not you receive a reply, it's very likely that your submission has been viewed and noted by someone on the team.\n\n***\n\n#####Rules:\n1. This is a no-meme subreddit. Only serious suggestions, reports, or questions allowed.\n\n2.  All content must be regarding the bot. Keep it on-topic please.\n\n3. Be respectful. We're all nice people here.\n\n***\n\n&amp;nbsp;\n\n####^(Please don't send a message before first submitting your post on the subreddit.)\n\n######**[Message us anyway.](https://www.reddit.com/message/compose?to=%2Fr%2FMemeInvestor_Bot)**",
			"got", v,
		)
	}
}

func TestGetCreated(t *testing.T) {
	sub := Subreddit{}
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	json.Unmarshal(data, &sub)
	if v := sub.GetCreated(); v != 1532014741.0 {
		t.Error(
			"For GetCreated()",
			"expected", 1532014741.0,
			"got", v,
		)
	}
}

func TestGetSubscribers(t *testing.T) {
	sub := Subreddit{}
	data, _ := ioutil.ReadFile("./tests/subreddit.json")
	json.Unmarshal(data, &sub)
	if v := sub.GetSubscribers(); v != 1339 {
		t.Error(
			"For GetSubscribers()",
			"expected", 1339,
			"got", v,
		)
	}
}
