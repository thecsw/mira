package goraw

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// When we initialize the Reddit instance,
// automatically start a goroutine that will
// update the token every 45 minutes. The
// auto_refresh should not be accessible to
// the end user as it is an internal method
func Init(c Credentials) *Reddit {
	auth, _ := Authenticate(&c)
	go auth.auto_refresh()
	return auth
}

// This goroutine reauthenticates the user
// every 45 minutes. It should be run with the go
// statement
func (c *Reddit) auto_refresh() {
	for {
		time.Sleep(45 * time.Minute)
		c.update_creds()
	}
}

// Reauthenticate and updates the object itself
func (c *Reddit) update_creds() {
	temp, _ := Authenticate(&c.Creds)
	*c = *temp
}

// Returns an access_token acquired using the provided credentials
func Authenticate(c *Credentials) (*Reddit, error) {
	// URL to get access_token
	auth_url := RedditBase + "api/v1/access_token"

	// Define the data to send in the request
	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("username", c.Username)
	form.Add("password", c.Password)

	// Encode the Authorization Header
	raw := c.ClientId + ":" + c.ClientSecret
	encoded := b64.StdEncoding.EncodeToString([]byte(raw))

	// Create a request to allow customised headers
	r, _ := http.NewRequest("POST", auth_url, strings.NewReader(form.Encode()))
	// Customise request headers
	r.Header.Set("User-Agent", c.UserAgent)
	r.Header.Set("Authorization", "Basic "+encoded)

	// Create client
	client := &http.Client{}

	// Run the request
	response, err := client.Do(r)
	defer response.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)

	auth := Reddit{}
	json.Unmarshal(buf.Bytes(), &auth)
	auth.Creds = *c
	return &auth, err
}
