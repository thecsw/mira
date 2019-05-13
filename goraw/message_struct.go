package goraw

type Listing struct {
	Data ListingData
}

type ListingData struct {
	Modhash string
	Dist float64
	Children []ListingDataChildren
	After string
	Before string
}

type ListingDataChildren struct {
	Kind string
	Data ListingDataChildrenData
}

type ListingDataChildrenData struct {
	First_message string
	First_message_name string
	Subreddit string
	Likes string
	Replies string
	Id string
	Subject string
	Was_comment bool
	Score float64
	Author string
	Num_comments float64
	Parent_id string
	Subreddit_name_prefixed string
	New bool
	Body string
	Link_title string
	Dest string
	Body_html string
	Name string
	Created float64
	Created_utc float64
	Context string
	Distinguished string
}
