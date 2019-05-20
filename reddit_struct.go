package mira

import (
	"time"
)

type Reddit struct {
	Token    string  `json:"access_token"`
	Duration float64 `json:"expires_in"`
	Creds    Credentials
	Stream   Streaming
}

type Streaming struct {
	CommentListInterval time.Duration
	PostListInterval    time.Duration
	PostListSlice       int
}
