package goraw

type Reddit struct {
	Token     string `json:"access_token"`
	Duration  int    `json:"expires_in"`
	Creds Credentials
}

type Credentials struct {
	ClientId     string
	ClientSecret string
	Username     string
	Password     string
	UserAgent    string
}
