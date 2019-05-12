package main

import (
	`./goraw`
	`fmt`
	`time`
)

func main() {
	cred := goraw.Credentials {
		"PMQYCRTBZq6qHw",
		"oLWXPkWF80zF8KnejyGjvZB_6VE",
		"HiveWriting_bot",
		"0p3244wGerMIYDZD",
		"HiveWriting_bot Ubuntu 16.04 (Reddit Hive Mind is writing a poem)",
	}
	r := goraw.Init(cred)
	fmt.Println(r)
	r.UpdateCreds()
	fmt.Println(r)
	go r.AutoRefresh() // Enables auto refreshing of the tokens
	time.Sleep(5 * time.Second)
//	fmt.Println(r.Me())
//	fmt.Println(r.GetUser("thecsw"))
}
