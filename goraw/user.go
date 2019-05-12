package goraw

import (
	"./redditor"
	"bytes"
	"encoding/json"
	"net/http"
)

// We have to make a new structure for Me
// I can do it tomorrow. It will be similar
// to redditor.Redditor but for some reason
// it should be different
func (c *Reddit) Me() redditor.Redditor {
	target := Authed_base + "api/v1/me"
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

func (c *Reddit) GetUser(name string) redditor.Redditor {
	target := Authed_base + "user/" + name + "/about"
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
