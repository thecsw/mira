package models

func (rl ReportListingChild) GetKind() string               { return rl.Kind }
func (rl *ReportListing) GetChildren() []ReportListingChild { return rl.Data.Children }
func (rl ReportListingChild) GetId() string {
	return rl.Data.Name
}
func (rl *ReportListingChild) GetUserReports() []UserReport {
	reports := make([]UserReport, 0)
	for i := range rl.Data.UserReports {
		report := UserReport{}
		report.Reason = rl.Data.UserReports[i].([]any)[0].(string)
		report.NumOfReports = rl.Data.UserReports[i].([]any)[1].(float64)
		report.SnoozeStatus = rl.Data.UserReports[i].([]any)[2].(bool)
		report.CanSnooze = rl.Data.UserReports[i].([]any)[3].(bool)
		reports = append(reports, report)
	}
	return reports
}
