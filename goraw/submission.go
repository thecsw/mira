package goraw

func (s *Submission) GetId() string {
	return s.Json.Data.Name
}
