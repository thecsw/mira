package main

import (
	"github.com/thecsw/mira/v4"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))

	r.Redditor("myuser").Compose("mytitle", "mytext")
}
