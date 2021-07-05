package mira

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Credentials struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	UserAgent    string
}

// Authenticate returns *Reddit object that has been authed
func Authenticate(c *Credentials) (*Reddit, error) {
	// URL to get access_token
	authURL := RedditBase + "api/v1/access_token"

	// Define the data to send in the request
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("username", c.Username)
	form.Add("password", c.Password)

	// Encode the Authorization Header
	raw := c.ClientID + ":" + c.ClientSecret
	encoded := b64.StdEncoding.EncodeToString([]byte(raw))

	// Create a request to allow customised headers
	r, err := http.NewRequest("POST", authURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	// Customise request headers
	r.Header.Set("User-Agent", c.UserAgent)
	r.Header.Set("Authorization", "Basic "+encoded)

	// Create client
	client := &http.Client{}

	// Run the request
	response, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	data := buf.Bytes()
	if err := findRedditError(data); err != nil {
		return nil, err
	}

	auth := Reddit{}
	auth.Chain = make(chan *ChainVals, 32)
	json.Unmarshal(data, &auth)
	auth.Creds = *c
	return &auth, nil
}

// This goroutine reauthenticates the user
// every 45 minutes. It should be run with the go
// statement
func (c *Reddit) autoRefresh() {
	for {
		time.Sleep(45 * time.Minute)
		c.updateCredentials()
	}
}

// Reauthenticate and updates the object itself
func (c *Reddit) updateCredentials() {
	temp, _ := Authenticate(&c.Creds)
	// Just updated the token
	c.Token = temp.Token
}

// SetDefault sets all default values
func (c *Reddit) SetDefault() {
	c.Stream = Streaming{
		CommentListInterval: 8,
		PostListInterval:    10,
		PostListSlice:       25,
	}
	c.Values = RedditVals{
		GetSubmissionFromCommentTries: 32,
	}
}

// SetClient sets mira's *http.Client to make requests
func (c *Reddit) SetClient(client *http.Client) {
	c.Client = client
}
