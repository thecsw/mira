package mira

import "time"

func (p *PostListing) GetChildren() []PostListingChild {
	return p.Data.Children
}

func (plc *PostListingChild) GetAge() int64 {
	return time.Now().UTC().Unix() - plc.Data.CreatedUtc
}

func (plc *PostListingChild) GetTimeCreated() int64 {
	return plc.Data.CreatedUtc
}
