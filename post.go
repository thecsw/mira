package mira

import "time"

func (p *PostListing) GetChildren() []PostListingChild {
	return p.Data.Children
}

func (plc *PostListingChild) GetAge() float64 {
	return float64(time.Now().UTC().Unix()) - plc.Data.CreatedUtc
}

func (plc *PostListingChild) GetTimeCreated() float64 {
	return plc.Data.CreatedUtc
}

func (plc *PostListingChild) GetSubreddit() string {
	return plc.Data.Subreddit
}

func (plc *PostListingChild) GetSubredditId() string {
	return plc.Data.SubredditId
}

func (plc *PostListingChild) GetTitle() string {
	return plc.Data.Title
}

func (plc *PostListingChild) GetId() string {
	return plc.Data.Name
}

func (plc *PostListingChild) GetAuthor() string {
	return plc.Data.Author
}

func (plc *PostListingChild) GetAuthorId() string {
	return plc.Data.AuthorFullname
}

func (plc *PostListingChild) GetCreated() float64 {
	return plc.Data.Created
}

func (plc *PostListingChild) GetUps() float64 {
	return plc.Data.Ups
}

func (plc *PostListingChild) GetDowns() float64 {
	return plc.Data.Downs
}

func (plc *PostListingChild) GetScore() float64 {
	return plc.Data.Score
}

func (plc *PostListingChild) GetText() string {
	return plc.Data.Selftext
}

func (plc *PostListingChild) GetAuthorFlair() string {
	return plc.Data.AuthorFlairText
}

func (plc *PostListingChild) GetPermalink() string {
	return plc.Data.Permalink
}

func (plc *PostListingChild) GetUrl() string {
	return plc.Data.Url
}

func (plc *PostListingChild) GetCommentCount() float64 {
	return plc.Data.NumComments
}

func (plc *PostListingChild) GetCrosspostCount() float64 {
	return plc.Data.NumCrossposts
}
