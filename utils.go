package mira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"slices"
	"strings"
)

// Short runes to variablize our types.
const (
	submissionType = "s"
	subredditType  = "b"
	commentType    = "c"
	redditorType   = "r"
	meType         = "m"
)

func (c *Reddit) checkType(rtype ...string) (string, string, error) {
	name, ttype := c.getQueue()
	if name == "" {
		return "", "", fmt.Errorf("identifier is empty")
	}
	if !findElem(ttype, rtype) {
		return "", "", fmt.Errorf(
			"the passed type is not a valid type for this call | expected: %s",
			strings.Join(rtype, ", "))
	}
	return name, ttype, nil
}

func (c *Reddit) addQueue(name string, ttype string) *Reddit {
	c.Chain <- &ChainVals{Name: name, Type: ttype}
	return c
}

func (c *Reddit) getQueue() (string, string) {
	if len(c.Chain) < 1 {
		return "", ""
	}
	temp := <-c.Chain
	return temp.Name, temp.Type
}

func findElem(elem string, arr []string) bool {
	return slices.Contains(arr, elem)
}

// RedditErr is a struct to store reddit error messages.
type RedditErr struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func findRedditError(data []byte) error {
	object := &RedditErr{}
	json.Unmarshal(data, object)
	if object.Message != "" || object.Error != "" {
		return fmt.Errorf("%s | error code: %s", object.Message, object.Error)
	}
	return nil
}

// ReadCredsFromFile reads mira credentials from a given file path
func ReadCredsFromFile(file string) Credentials {
	// Declare all regexes
	ClientID, _ := regexp.Compile(`CLIENT_ID\s*=\s*(.+)`)
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
		ClientID.FindStringSubmatch(s)[1],
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
