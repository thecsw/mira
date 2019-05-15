package mira

type Reddit struct {
	Token    string  `json:"access_token"`
	Duration float64 `json:"expires_in"`
	Creds    Credentials
}
