package mira

func (l *CommentListing) GetMessages() []CommentListingDataChildren {
	return l.Data.Children
}

func (ldc *CommentListingDataChildren) GetId() string {
	return ldc.Data.Name
}

func (ldc *CommentListingDataChildren) IsComment() bool {
	return ldc.Kind == "t1"
}

func (ldc *CommentListingDataChildren) IsCommentReply() bool {
	return ldc.Data.Subject == "comment reply"
}
