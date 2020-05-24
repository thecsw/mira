package mira

import (
	"io/ioutil"
	"os"
	"regexp"
)

// ReadCredsFromFile reads mira credentials from a given file path
func ReadCredsFromFile(file string) Credentials {
	// Declare all regexes
	ClientId, _ := regexp.Compile(`CLIENT_ID\s*=\s*(.+)`)
	ClientSecret, _ := regexp.Compile(`CLIENT_SECRET\s*=\s*(.+)`)
	Username, _ := regexp.Compile(`USERNAME\s*=\s*(.+)`)
	Password, _ := regexp.Compile(`PASSWORD\s*=\s*(.+)`)
	UserAgent, _ := regexp.Compile(`USER_AGENT\s*=\s*(.+)`)
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return Credentials{}
	}
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

// ReadCredsFromEnv reads mira credentials from environment
func ReadCredsFromEnv() Credentials {
	return Credentials{
		os.Getenv("BOT_CLIENT_ID"),
		os.Getenv("BOT_CLIENT_SECRET"),
		os.Getenv("BOT_USERNAME"),
		os.Getenv("BOT_PASSWORD"),
		os.Getenv("BOT_USER_AGENT"),
	}
}
