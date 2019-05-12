package goraw

import (
	`./auth`
	`./redditor`
	`net/http`
	`bytes`
	`encoding/json`
	`fmt`
)

type Reddit struct {
	Token     string `json:"access_token"`
	Duration  int    `json:"expires_in"`
	UserAgent string
}

func Init(id, sec, user, pass, agent string) (Reddit) {
	cred := auth.Credentials {id,sec,user,pass,agent}
	auth, _ := auth.Authenticate(&cred)
	r := Reddit{
		auth.Token,
		auth.Duration,
		auth.UserAgent,
	}
	return r
}

func (c* Reddit) Me() redditor.Redditor {
	target := auth.Authed_base + "api/v1/me"
	r, _ := http.NewRequest("GET", target, nil)
	r.Header.Set("User-Agent", c.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	user := redditor.Redditor{}
	json.Unmarshal(buf.Bytes(), &user)
	fmt.Println(string(buf.Bytes()))
	return user
}

func (c* Reddit) GetUser(name string) redditor.Redditor {
	target := auth.Authed_base + "user/" + name + "/about"
	r, _ := http.NewRequest("GET", target, nil)
	r.Header.Set("User-Agent", c.UserAgent)
	r.Header.Set("Authorization", "bearer "+c.Token)
	client := &http.Client{}
	response, _ := client.Do(r)
	defer response.Body.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	user := redditor.Redditor{}
	json.Unmarshal(buf.Bytes(), &user)
	return user
}
