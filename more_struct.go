package mira

type MoreListing struct {
	Kind string          `json:"kind"`
	Data MoreListingData `json:"data"`
}

type MoreListingData struct {
	Count    float64  `json:"count"`
	Name     string   `json:"name"`
	Id       string   `json:"id"`
	ParentId string   `json:"parent_id"`
	Depth    float64  `json:"depth"`
	Children []string `json:"children"`
}
