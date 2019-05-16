package main

import (
	"github.com/thecsw/mira"
	"fmt"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.conf"))

	r.Compose("thecsw", "my subject", "hello, world")
	// or
	user, _ := r.GetUser("thecsw")
	r.Compose(user.GetName(), "my subject", "hello, world")
}
