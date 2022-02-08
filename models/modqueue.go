package models

func (mql ModQueueListingChild) GetKind() string { return mql.Kind }
func (mql ModQueueListingChild) GetId() string {
	if mql.GetKind() == "t1" {
		return mql.Data.(CommentListingDataChildrenData).Name
	} else {
		return mql.Data.(PostListingChildData).Name
	}
}
func (mql *ModQueueListing) GetChildren() []ModQueueListingChild { return mql.Data.Children }
