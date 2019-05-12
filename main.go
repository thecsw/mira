package main

import (
	"fmt"

	"./goraw"
)

func main() {
	cred := goraw.Credentials{
		"PMQYCRTBZq6qHw",
		"oLWXPkWF80zF8KnejyGjvZB_6VE",
		"HiveWriting_bot",
		"0p3244wGerMIYDZD",
		"HiveWriting_bot Ubuntu 16.04 (Reddit Hive Mind is writing a poem)",
	}
	// It will handle aute refreshing the tokens
	r := goraw.Init(cred)
	fmt.Println(r)

	fmt.Println(r.Me())
	fmt.Println(r.GetUser("thecsw"))
}
