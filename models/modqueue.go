package models

func (mql ModQueueListingChild) GetKind() string { return mql.Kind }
func (mql ModQueueListingChild) GetId() string {
	return mql.Data.Name
}
func (mql *ModQueueListing) GetChildren() []ModQueueListingChild { return mql.Data.Children }
