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
