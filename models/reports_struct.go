package models

type ReportListing struct {
	Kind string            `json:"kind"`
	Data ReportListingData `json:"data"`
}

type ReportListingData struct {
	Modhash  string               `json:"modhash"`
	Dist     float64              `json:"dist"`
	Children []ReportListingChild `json:"children"`
}

type ReportListingChild struct {
	Kind string               `json:"kind"`
	Data PostListingChildData `json:"data"`
}
