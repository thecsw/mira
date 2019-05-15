package mira

// Get the ID of the submission in form "t3_..."
func (s *Submission) GetId() string {
	return s.Json.Data.Name
}

//
func (s *Submission) GetDraftsCount() float64 {
	return s.Json.Data.DraftsCount
}

func (s *Submission) GetUrl() string {
	return s.Json.Data.Url
}
