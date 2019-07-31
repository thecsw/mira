package mira

import (
	"time"

	"github.com/thecsw/mira/models"
)

// // c is the channel with all unread messages
// // stop is the channel to stop the stream. Do stop <- true to stop the loop
func (r *Reddit) StreamCommentReplies() (<-chan models.Comment, chan bool) {
	c := make(chan models.Comment, 25)
	stop := make(chan bool, 1)
	go func() {
		for {
			stop <- false
			un, _ := r.ListUnreadMessages()
			for _, v := range un {
				// Only process comment replies and
				// mark them as read.
				if v.IsCommentReply() {
					c <- v
					r.ReadMessage(v.GetId())
				}
			}
			time.Sleep(r.Stream.CommentListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop
}

// // c is the channel with all comments
// // stop is the channel to stop the stream. Do stop <- true to stop the loop
func (r *Reddit) StreamComments() (<-chan models.Comment, chan bool, error) {
	err := r.checkType("subreddit", "redditor")
	if err != nil {
		return nil, nil, err
	}
	switch r.Chain.Type {
	case "subreddit":
		return r.streamSubredditComments()
	case "redditor":
		return r.streamRedditorComments()
	}
	return nil, nil, nil
}

func (r *Reddit) streamSubredditComments() (<-chan models.Comment, chan bool, error) {
	c := make(chan models.Comment, 25)
	stop := make(chan bool, 1)
	anchor, err := r.Subreddit(r.Chain.Name).Comments("new", "hour", 1)
	if err != nil {
		return nil, nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			stop <- false
			un, _ := r.Subreddit(r.Chain.Name).CommentsAfter("new", last, 25)
			for _, v := range un {
				c <- v
			}
			if len(un) > 0 {
				last = un[0].GetId()
			}
			time.Sleep(r.Stream.CommentListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop, nil
}

func (r *Reddit) streamRedditorComments() (<-chan models.Comment, chan bool, error) {
	c := make(chan models.Comment, 25)
	stop := make(chan bool, 1)
	anchor, err := r.Redditor(r.Chain.Name).Comments("new", "hour", 1)
	if err != nil {
		return nil, nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			stop <- false
			un, _ := r.Redditor(r.Chain.Name).CommentsAfter("new", last, 25)
			for _, v := range un {
				c <- v
			}
			if len(un) > 0 {
				last = un[0].GetId()
			}
			time.Sleep(r.Stream.CommentListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop, nil
}

func (r *Reddit) StreamSubmissions() (<-chan models.PostListingChild, chan bool, error) {
	err := r.checkType("subreddit", "redditor")
	if err != nil {
		return nil, nil, err
	}
	switch r.Chain.Type {
	case "subreddit":
		return r.streamSubredditSubmissions()
	case "redditor":
		return r.streamRedditorSubmissions()
	}
	return nil, nil, nil
}

func (r *Reddit) streamSubredditSubmissions() (<-chan models.PostListingChild, chan bool, error) {
	c := make(chan models.PostListingChild, 25)
	stop := make(chan bool, 1)
	anchor, err := r.Subreddit(r.Chain.Name).Submissions("new", "hour", 1)
	if err != nil {
		return nil, nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			stop <- false
			new, _ := r.Subreddit(r.Chain.Name).SubmissionsAfter(last, r.Stream.PostListSlice)
			if len(new) > 0 {
				last = new[0].GetId()
			}
			for i := range new {
				c <- new[len(new)-i-1]
			}
			time.Sleep(r.Stream.PostListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop, nil
}

func (r *Reddit) streamRedditorSubmissions() (<-chan models.PostListingChild, chan bool, error) {
	c := make(chan models.PostListingChild, 25)
	stop := make(chan bool, 1)
	anchor, err := r.Redditor(r.Chain.Name).Submissions("new", "hour", 1)
	if err != nil {
		return nil, nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			stop <- false
			new, _ := r.Redditor(r.Chain.Name).SubmissionsAfter(last, r.Stream.PostListSlice)
			if len(new) > 0 {
				last = new[0].GetId()
			}
			for i := range new {
				c <- new[len(new)-i-1]
			}
			time.Sleep(r.Stream.PostListInterval * time.Second)
			if <-stop {
				return
			}
		}
	}()
	return c, stop, nil
}
