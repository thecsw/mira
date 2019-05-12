package goraw

import (
	"io/ioutil"
	"regexp"
)

func ReadCredsFromFile(file string) Credentials {
	pattern := `CLIENT_ID = (.+)\nCLIENT_SECRET = (.+)\nUSERNAME = (.+)\nPASSWORD = (.+)\nUSER_AGENT = (.+)`
	r, _ := regexp.Compile(pattern)
	data, _ := ioutil.ReadFile(file)
	s := string(data)
	creds := Credentials{
		r.FindStringSubmatch(s)[1],
		r.FindStringSubmatch(s)[2],
		r.FindStringSubmatch(s)[3],
		r.FindStringSubmatch(s)[4],
		r.FindStringSubmatch(s)[5],
	}
	return creds
}
