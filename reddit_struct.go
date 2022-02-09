package mira

import (
	"net/http"
	"time"
)

// Reddit is the main mira struct that practically
// does everything
type Reddit struct {
	Token    string  `json:"access_token"`
	Duration float64 `json:"expires_in"`
	Creds    Credentials
	Chain    chan *ChainVals
	Stream   Streaming
	Values   RedditVals
	Client   *http.Client
}

// Streaming is used for some durations on how frequently
// do we listen to comments/submissions
type Streaming struct {
	CommentListInterval time.Duration
	PostListInterval    time.Duration
	PostListSlice       int
	ReportsInterval     time.Duration
	ModQueueInterval    time.Duration
}

// RedditVals is just some values to backoff
type RedditVals struct {
	GetSubmissionFromCommentTries int
}

// ChainVals is our queue values
type ChainVals struct {
	Name string
	Type string
}
