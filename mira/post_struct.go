package mira

type PostListing struct {
	Kind    string          `json:"kind"`
	Data    PostListingData `json:"data"`
}

type PostListingData struct {
	After    string             `json:"after"`
	Children []PostListingChild `json:"children"`
}

type PostListingChild struct{
	Kind string `json:"kind"`
	Data Post   `json:"post"`
}

type Post struct {
	Subreddit      string `json:"subreddit"`
	AuthorFullname string `json:"author_fullname"`
	HideScore      bool   `json:"hide_score"`
	Downvotes      int64  `json:"downs"`
	Upvotes        int64  `json:"ups"`
	Name           string `json:"name"`
	Created        int64  `json:"created"`
	CreatedUtc     int64  `json:"created_utc"`
	SubredditId    string `json:"subreddit_id"`
	Author         string `json:"author"`
	Id             string `json:"id"`
	NumComments    string `json:"num_comments"`
	Url            string `json:"url"`
	Permalink      string `json:"permalink"`
}