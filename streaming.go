package mira

import (
	"time"

	"github.com/thecsw/mira/models"
)

// StreamCommentReplies streams comment replies
// c is the channel with all unread messages
func (c *Reddit) StreamCommentReplies() <-chan models.Comment {
	ret := make(chan models.Comment, 100)
	go func() {
		for {
			un, _ := c.Me().ListUnreadMessages()
			for _, v := range un {
				if v.IsCommentReply() {
					// Only process comment replies and
					// mark them as read.
					ret <- v
					// You can read the message with
					c.Me().ReadMessage(v.GetId())
				}
			}
			time.Sleep(c.Stream.CommentListInterval * time.Second)
		}
	}()
	return ret
}

// StreamMentions streams recent mentions
// c is the channel with all unread messages
func (c *Reddit) StreamMentions() <-chan models.Comment {
	ret := make(chan models.Comment, 100)
	go func() {
		for {
			un, _ := c.Me().ListUnreadMessages()
			for _, v := range un {
				if v.IsMention() {
					// Only process comment replies and
					// mark them as read.
					ret <- v
					// You can read the message with
					c.Me().ReadMessage(v.GetId())
				}
			}
			time.Sleep(c.Stream.CommentListInterval * time.Second)
		}
	}()
	return ret
}

// StreamComments streams comments from a redditor or a subreddit
// c is the channel with all comments
func (c *Reddit) StreamComments() (<-chan models.Comment, error) {
	name, ttype, err := c.checkType(subredditType, redditorType)
	if err != nil {
		return nil, err
	}
	switch ttype {
	case subredditType:
		return c.streamSubredditComments(name)
	case redditorType:
		return c.streamRedditorComments(name)
	}
	return nil, nil
}

// StreamSubmissions streams submissions from a redditor or a subreddit
// c is the channel with all submissions.
func (c *Reddit) StreamSubmissions() (<-chan models.PostListingChild, error) {
	name, ttype, err := c.checkType(subredditType, redditorType)
	if err != nil {
		return nil, err
	}
	switch ttype {
	case subredditType:
		return c.streamSubredditSubmissions(name)
	case redditorType:
		return c.streamRedditorSubmissions(name)
	}
	return nil, nil
}

// StreamModQueue streams modqueue entries from a subreddit
// c is the channel with all modqueue entries.
func (c *Reddit) StreamModQueue() (<-chan models.ModQueueListingChild, error) {
	name, ttype, err := c.checkType(subredditType)
	if err != nil {
		return nil, err
	}
	switch ttype {
	case subredditType:
		return c.streamSubredditModQueue(name)
	}
	return nil, nil
}

func (c *Reddit) streamSubredditModQueue(subreddit string) (<-chan models.ModQueueListingChild, error) {
	ret := make(chan models.ModQueueListingChild, 100)
	anchor, err := c.Subreddit(subreddit).ModQueue(1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			new, _ := c.Subreddit(subreddit).ModQueueAfter(last, c.Stream.PostListSlice)
			if len(new) < 1 {
				time.Sleep(c.Stream.ModQueueInterval * time.Second)
				continue
			}
			last = new[0].GetId()
			for i := range new {
				ret <- new[len(new)-i-1]
			}
			time.Sleep(c.Stream.ModQueueInterval * time.Second)
		}
	}()
	return ret, nil
}

// StreamReports streams reports entries from a subreddit
// c is the channel with all report entries.
func (c *Reddit) StreamReports() (<-chan models.ReportListingChild, error) {
	name, ttype, err := c.checkType(subredditType)
	if err != nil {
		return nil, err
	}
	switch ttype {
	case subredditType:
		return c.streamSubredditReports(name)
	}
	return nil, nil
}

func (c *Reddit) streamSubredditReports(subreddit string) (<-chan models.ReportListingChild, error) {
	ret := make(chan models.ReportListingChild, 100)
	anchor, err := c.Subreddit(subreddit).Reports(1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			new, _ := c.Subreddit(subreddit).ReportsAfter(last, c.Stream.PostListSlice)
			if len(new) < 1 {
				time.Sleep(c.Stream.ReportsInterval * time.Second)
				continue
			}
			last = new[0].GetId()
			for i := range new {
				ret <- new[len(new)-i-1]
			}
			time.Sleep(c.Stream.ReportsInterval * time.Second)
		}
	}()
	return ret, nil
}

func (c *Reddit) streamSubredditComments(subreddit string) (<-chan models.Comment, error) {
	ret := make(chan models.Comment, 100)
	anchor, err := c.Subreddit(subreddit).Comments(New, Hour, 1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			un, _ := c.Subreddit(subreddit).CommentsAfter(New, last, 100)
			if len(un) < 1 {
				time.Sleep(c.Stream.CommentListInterval * time.Second)
				continue
			}
			last = un[0].GetId()
			for _, v := range un {
				ret <- v
			}
			time.Sleep(c.Stream.CommentListInterval * time.Second)
		}
	}()
	return ret, nil
}

func (c *Reddit) streamRedditorComments(redditor string) (<-chan models.Comment, error) {
	ret := make(chan models.Comment, 100)
	anchor, err := c.Redditor(redditor).Comments(New, Hour, 1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			un, _ := c.Redditor(redditor).CommentsAfter(New, last, 100)
			if len(un) < 1 {
				time.Sleep(c.Stream.CommentListInterval * time.Second)
				continue
			}
			last = un[0].GetId()
			for _, v := range un {
				ret <- v
			}
			time.Sleep(c.Stream.CommentListInterval * time.Second)
		}
	}()
	return ret, nil
}

func (c *Reddit) streamSubredditSubmissions(subreddit string) (<-chan models.PostListingChild, error) {
	ret := make(chan models.PostListingChild, 100)
	anchor, err := c.Subreddit(subreddit).Submissions(New, Hour, 1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			new, _ := c.Subreddit(subreddit).SubmissionsAfter(last, c.Stream.PostListSlice)
			if len(new) < 1 {
				time.Sleep(c.Stream.PostListInterval * time.Second)
				continue
			}
			last = new[0].GetId()
			for i := range new {
				ret <- new[len(new)-i-1]
			}
			time.Sleep(c.Stream.PostListInterval * time.Second)
		}
	}()
	return ret, nil
}

func (c *Reddit) streamRedditorSubmissions(redditor string) (<-chan models.PostListingChild, error) {
	ret := make(chan models.PostListingChild, 100)
	anchor, err := c.Redditor(redditor).Submissions(New, Hour, 1)
	if err != nil {
		return nil, err
	}
	last := ""
	if len(anchor) > 0 {
		last = anchor[0].GetId()
	}
	go func() {
		for {
			new, _ := c.Redditor(redditor).SubmissionsAfter(last, c.Stream.PostListSlice)
			if len(new) < 1 {
				time.Sleep(c.Stream.PostListInterval * time.Second)
				continue
			}
			last = new[0].GetId()
			for i := range new {
				ret <- new[len(new)-i-1]
			}
			time.Sleep(c.Stream.PostListInterval * time.Second)
		}
	}()
	return ret, nil
}
