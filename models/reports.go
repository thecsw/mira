package models

func (rl ReportListingChild) GetKind() string { return rl.Kind }
func (rl ReportListingChild) GetId() string {
	if rl.GetKind() == "t1" {
		return rl.Data.(CommentListingDataChildrenData).Name
	} else {
		return rl.Data.(PostListingChildData).Name
	}
}
func (rl *ReportListing) GetChildren() []ReportListingChild { return rl.Data.Children }
