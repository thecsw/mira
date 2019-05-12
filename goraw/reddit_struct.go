package goraw

type Reddit struct {
	Token    string `json:"access_token"`
	Duration int    `json:"expires_in"`
	Creds    Credentials
}
