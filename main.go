package main

import (
	`./src`
	`fmt`
)

func main() {
	r := goraw.Init(
		"PMQYCRTBZq6qHw",
		"oLWXPkWF80zF8KnejyGjvZB_6VE",
		"HiveWriting_bot",
		"0p3244wGerMIYDZD",
		"HiveWriting_bot Ubuntu 16.04 (Reddit Hive Mind is writing a poem)",
	)
	fmt.Println(r.Me())
	fmt.Println(r.GetUser("thecsw"))
}
