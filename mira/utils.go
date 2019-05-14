package mira

import (
	"io/ioutil"
	"regexp"
)

func ReadCredsFromFile(file string) Credentials {
	// Declare all regexes
	ClientId, _ := regexp.Compile(`CLIENT_ID\s*=\s*(.+)`)
	ClientSecret, _ := regexp.Compile(`CLIENT_SECRET\s*=\s*(.+)`)
	Username, _ := regexp.Compile(`USERNAME\s*=\s*(.+)`)
	Password, _ := regexp.Compile(`PASSWORD\s*=\s*(.+)`)
	UserAgent, _ := regexp.Compile(`USER_AGENT\s*=\s*(.+)`)
	data, _ := ioutil.ReadFile(file)
	s := string(data)
	creds := Credentials{
		ClientId.FindStringSubmatch(s)[1],
		ClientSecret.FindStringSubmatch(s)[1],
		Username.FindStringSubmatch(s)[1],
		Password.FindStringSubmatch(s)[1],
		UserAgent.FindStringSubmatch(s)[1],
	}
	return creds
}
