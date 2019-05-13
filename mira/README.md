# MIRA or Meme Investor Reddit Api

MIRA is a Reddit Api Wrapper written in beautiful Go. This is a subproject
of a bigger project MemeInvestor_bot. 

IT is super simple to use the bot as we also provide you
with simple but fully extensive structs. We utilize the 
best features of Go, such as closures, channels, goroutines, garbage collection, etc.

Below is an example on how to make a simple bot that 
listens to a stream of comment replies and replies.

``` go
package main

import (
	"github.com/thecsw/mira"
	"fmt"
)

func main() {
	r, _ := mira.Init(mira.ReadCredsFromFile("login.cong"))
	c, _ := r.StreamCommentReplies()
	for {
		msg := <- c
		r.Reply(msg.GetId(), "I got your message!")
	}
}
```

## Config file

The config file structure is very simple:

```
CLIENT_ID =
CLIENT_SECRET =
USERNAME =
PASSWORD =
USER_AGENT =
```


