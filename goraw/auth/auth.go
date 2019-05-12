package auth

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

type Authorization struct {
	Token     string `json:"access_token"`
	Duration  int    `json:"expires_in"`
	UserAgent string
}

type Credentials struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	UserAgent    string
}

// Base is the basic reddit base URL, authed is the base URL for use once authenticated
var Base string = "https://www.reddit.com/"
var Authed_base string = "https://oauth.reddit.com/"

// Returns an access_token acquired using the provided credentials
func Authenticate(c *Credentials) (*Authorization, error) {
	// URL to get access_token
	auth_url := Base + "api/v1/access_token"

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

	var auth Authorization
	json.Unmarshal(buf.Bytes(), &auth)
	auth.UserAgent = c.UserAgent

	return &auth, err
}
