package mira

import (
	"time"

	"github.com/thecsw/mira/models"
)

// c is the channel with all unread messages
// stop is the channel to stop the stream. Do stop <- true to stop the loop
func (r *Reddit) StreamCommentReplies() <-chan models.Comment {
	c := make(chan models.Comment, 100)
	go func() {
		for {
			un, _ := r.Me().ListUnreadMessages()
			for _, v := range un {
				if v.IsCommentReply() {
					// Only process comment replies and
					// mark them as read.
					c <- v
					// You can read the message with
					r.Me().ReadMessage(v.GetId())
				}
			}
			time.Sleep(r.Stream.CommentListInterval * time.Second)
		}
	}()
	return c
}

// c is the channel with all unread messages
// stop is the channel to stop the stream. Do stop <- true to stop the loop
func (r *Reddit) StreamMentions() <-chan models.Comment {
	c := make(chan models.Comment, 100)
	go func() {
		for {
			un, _ := r.Me().ListUnreadMessages()
			for _, v := range un {
				if v.IsMention() {
					// Only process comment replies and
					// mark them as read.
					c <- v
					// You can read the message with
					r.Me().ReadMessage(v.GetId())
				}
			}
			time.Sleep(r.Stream.CommentListInterval * time.Second)
		}
	}()
	return c
}

// c is the channel with all comments
// stop is the channel to stop the stream. Do stop <- true to stop the loop
func (r *Reddit) StreamComments() (<-chan models.Comment, error) {
	name, ttype, err := r.checkType("subreddit", "redditor")
	if err != nil {
		return nil, err
	}
	switch ttype {
	case "subreddit":
		return r.streamSubredditComments(name)
	case "redditor":
		return r.streamRedditorComments(name)
	}
	return nil, nil
}

func (r *Reddit) streamSubredditComments(subreddit string) (<-chan models.Comment, error) {
	c := make(chan models.Comment, 100)
	anchor, err := r.Subreddit(subreddit).Comments("new", "hour", 1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			un, _ := r.Subreddit(subreddit).CommentsAfter("new", last, 100)
			if len(un) < 0 {
				time.Sleep(r.Stream.CommentListInterval * time.Second)
				continue
			}
			last = un[0].GetId()
			for _, v := range un {
				c <- v
			}
			time.Sleep(r.Stream.CommentListInterval * time.Second)
		}
	}()
	return c, nil
}

func (r *Reddit) streamRedditorComments(redditor string) (<-chan models.Comment, error) {
	c := make(chan models.Comment, 100)
	anchor, err := r.Redditor(redditor).Comments("new", "hour", 1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			un, _ := r.Redditor(redditor).CommentsAfter("new", last, 100)
			if len(un) < 1 {
				time.Sleep(r.Stream.CommentListInterval * time.Second)
				continue
			}
			last = un[0].GetId()
			for _, v := range un {
				c <- v
			}
			time.Sleep(r.Stream.CommentListInterval * time.Second)
		}
	}()
	return c, nil
}

func (r *Reddit) StreamSubmissions() (<-chan models.PostListingChild, error) {
	name, ttype, err := r.checkType("subreddit", "redditor")
	if err != nil {
		return nil, err
	}
	switch ttype {
	case "subreddit":
		return r.streamSubredditSubmissions(name)
	case "redditor":
		return r.streamRedditorSubmissions(name)
	}
	return nil, nil
}

func (r *Reddit) streamSubredditSubmissions(subreddit string) (<-chan models.PostListingChild, error) {
	c := make(chan models.PostListingChild, 100)
	anchor, err := r.Subreddit(subreddit).Submissions("new", "hour", 1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			new, _ := r.Subreddit(subreddit).SubmissionsAfter(last, r.Stream.PostListSlice)
			if len(new) < 0 {
				time.Sleep(r.Stream.PostListInterval * time.Second)
				continue
			}
			last = new[0].GetId()
			for i := range new {
				c <- new[len(new)-i-1]
			}
			time.Sleep(r.Stream.PostListInterval * time.Second)
		}
	}()
	return c, nil
}

func (r *Reddit) streamRedditorSubmissions(redditor string) (<-chan models.PostListingChild, error) {
	c := make(chan models.PostListingChild, 100)
	anchor, err := r.Redditor(redditor).Submissions("new", "hour", 1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			new, _ := r.Redditor(redditor).SubmissionsAfter(last, r.Stream.PostListSlice)
			if len(new) < 0 {
				time.Sleep(r.Stream.PostListInterval * time.Second)
				continue
			}
			last = new[0].GetId()
			for i := range new {
				c <- new[len(new)-i-1]
			}
			time.Sleep(r.Stream.PostListInterval * time.Second)
		}
	}()
	return c, nil
}
