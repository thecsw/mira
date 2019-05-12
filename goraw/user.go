package goraw

import (
	"bytes"
	"encoding/json"
	"net/http"

	"./redditor"
)

func (c *Reddit) Me() redditor.Me {
	target := RedditApiMe
	r, _ := http.NewRequest("GET", target, nil)
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	user := redditor.Me{}
	json.Unmarshal(buf.Bytes(), &user)
	return user
}

func (c *Reddit) GetUser(name string) redditor.Redditor {
	target := RedditOauth + "user/" + name + "/about"
	r, _ := http.NewRequest("GET", target, nil)
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	user := redditor.Redditor{}
	json.Unmarshal(buf.Bytes(), &user)
	return user
}
