package goraw

import (
 	"bytes"
	"encoding/json"
	"net/http"
 	"net/url"
 	"strings"
)

func (c *Comment) Reply(rd *Reddit, text string) Comment {
	target := RedditOauth + "/api/comment"
	form := url.Values{}
	form.Add("text", text)
	form.Add("thing_id", c.Json.Data.Things[0].Data.Name)
	form.Add("api_type", "json")
	r, _ := http.NewRequest("POST", target, strings.NewReader(form.Encode()))
	r.Header.Set("User-Agent", rd.Creds.UserAgent)
	r.Header.Set("Authorization", "bearer "+rd.Token)
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	comment := Comment{}
	json.Unmarshal(buf.Bytes(), &comment)
	return comment
}
