package mira

type CommentListing struct {
	Kind string             `json:"kind"`
	Data CommentListingData `json:"data"`
}

type CommentListingData struct {
	Modhash  string    `json:"modhash"`
	Dist     float64   `json:"dist"`
	Children []Comment `json:"children"`
	After    string    `json:"after"`
	Before   string    `json:"before"`
}

type Comment struct {
	Kind string                         `json:"kind"`
	Data CommentListingDataChildrenData `json:"data"`
}

type CommentListingDataChildrenData struct {
	FirstMessage          string   `json:"first_message"`
	FirstMessageName      string   `json:"first_message_name"`
	Subreddit             string   `json:"subreddit"`
	Likes                 string   `json:"likes"`
	Replies               string   `json:"replies"`
	Id                    string   `json:"id"`
	Subject               string   `json:"subject"`
	WasComment            bool     `json:"was_comment"`
	Score                 float64  `json:"score"`
	Author                string   `json:"author"`
	NumComments           float64  `json:"num_comments"`
	ParentId              string   `json:"parent_id"`
	SubredditNamePrefixed string   `json:"subreddit_name_prefixed"`
	New                   bool     `json:"new"`
	Body                  string   `json:"body"`
	LinkTitle             string   `json:"link_title"`
	Dest                  string   `json:"dest"`
	BodyHtml              string   `json:"body_html"`
	Name                  string   `json:"name"`
	Created               float64  `json:"created"`
	Created_utc           float64  `json:"created_utc"`
	Context               string   `json:"context"`
	Distinguished         string   `json:"distinguished"`
	Children              []string `json:"children"`
}
