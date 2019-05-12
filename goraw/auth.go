package goraw

import (
	`bytes`
	b64 `encoding/base64`
	`encoding/json`
	`net/http`
	`net/url`
	`strings`
	`time`
)

func Init(c Credentials) *Reddit {
	auth, _ := Authenticate(&c)
	return auth
}

// This goroutine reauthenticates the user
// every hour. It should be run with the go
// statement
func (c* Reddit) AutoRefresh() {
	for ;; {
		time.Sleep(3600 * time.Second)
		c.UpdateCreds()
	}
}

// Reauthenticate and updates the object itself
func (c* Reddit) UpdateCreds() {
	temp, _ := Authenticate(&c.Creds)
	*c = *temp
}

// Returns an access_token acquired using the provided credentials
func Authenticate(c *Credentials) (*Reddit, error) {
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

	auth := Reddit{}
	json.Unmarshal(buf.Bytes(), &auth)
	auth.Creds = *c
	return &auth, err
}
