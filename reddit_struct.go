package mira

import (
	"net/http"
	"time"
)

// Reddit is the main mira struct that practically
// does everything
type Reddit struct {
	Token              string  `json:"access_token"`
	Duration           float64 `json:"expires_in"`
	Creds              Credentials
	Chain              chan *ChainVals
	Stream             Streaming
	Values             RedditVals
	Client             *http.Client
	RateLimitUsed      int     // The number of requests used in the current rate limit window
	RateLimitRemaining float64 // The number of requests left to use in the current rate limit window
	RateLimitReset     int     // The number of seconds left in the current rate limit window
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
