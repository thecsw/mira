package goraw

func (l* Listing) GetMessages() []ListingDataChildren {
	return l.Data.Children
}

func (ldc *ListingDataChildren) GetId() string {
	return ldc.Data.Name
}

func (ldc *ListingDataChildren) IsComment() bool {
	return ldc.Kind == "t1"
}

func (ldc *ListingDataChildren) IsCommentReply() bool {
	return ldc.Data.Subject == "comment reply"
}
