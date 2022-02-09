package mira

const (
	// RedditBase is the basic reddit base URL, authed is the base URL for use once authenticated
	RedditBase = "https://www.reddit.com/"
	// RedditOauth is the oauth url to pass the tokens to
	RedditOauth = "https://oauth.reddit.com"

	// JsonAPI sets the api type to json
	JsonAPI = "json"

	// Sorting options
	Hot           = "hot"
	New           = "new"
	Top           = "top"
	Rising        = "rising"
	Controversial = "controversial"
	Random        = "random"

	// Duration options
	Hour = "hour"
	Day  = "day"
	Week = "week"
	Year = "year"
	All  = "all"
)
