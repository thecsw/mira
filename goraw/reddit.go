package goraw

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

func (c *Reddit) Me() Me {
	target := RedditApiMe
	r, _ := http.NewRequest("GET", target, nil)
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	user := Me{}
	json.Unmarshal(buf.Bytes(), &user)
	return user
}

func (c *Reddit) GetUser(name string) Redditor {
	target := RedditOauth + "/user/" + name + "/about"
	r, _ := http.NewRequest("GET", target, nil)
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	user := Redditor{}
	json.Unmarshal(buf.Bytes(), &user)
	return user
}

func (c *Reddit) Submit(title, text, sr string) Submission {
	target := RedditOauth + "/api/submit"
	form := url.Values{}
	form.Add("title", title)
	form.Add("sr", sr)
	form.Add("text", text)
	form.Add("kind", "self")
	form.Add("resubmit", "true")
	form.Add("api_type", "json")
	r, _ := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	r.Header.Set("User-Agent", c.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	post := Submission{}
	json.Unmarshal(buf.Bytes(), &post)
	return post
}