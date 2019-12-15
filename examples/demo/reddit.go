package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/thecsw/mira"
)

func main() {
	r, err := mira.Init(mira.ReadCredsFromFile("login.conf"))
	check(err)
	log.Infoln("Connected to reddit!")

	me, err := r.Me().Info()
	check(err)
	log.Infoln("My reddit name is", me.GetAuthor())

	log.Infoln("Let's listen to some hot r/all stuff!")
	posts, err := r.Subreddit("all").Submissions("hot", "day", 5)
	check(err)
	for i, v := range posts {
		log.Infof("%d | %s by %s (%0.f upvotes)",
			i, v.GetId(), v.GetAuthor(), v.GetUps())
	}

	log.Infoln("That's it!")
}

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}
