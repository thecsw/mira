package mira

type Submission struct {
	Json SubmissionJson `json:"json"`
}

type SubmissionJson struct {
	Errors []string           `json:"errors"`
	Data   SubmissionJsonData `json:"data"`
}

type SubmissionJsonData struct {
	Url         string  `json:"url"`
	DraftsCount float64 `json:"drafts_count"`
	Id          string  `json:"id"`
	Name        string  `json:"name"`
}
