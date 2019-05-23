package mira

func (l *CommentListing) GetChildren() []CommentListingDataChildren {
	return l.Data.Children
}

func (ldc *CommentListingDataChildren) GetId() string {
	return ldc.Data.Name
}

func (ldc *CommentListingDataChildren) GetParentId() string {
	return ldc.Data.ParentId
}

func (ldc *CommentListingDataChildren) IsRoot() bool {
	return string(ldc.Data.ParentId[1]) == "3"
}

func (ldc *CommentListingDataChildren) GetBody() string {
	return ldc.Data.Body
}

func (ldc *CommentListingDataChildren) GetSubreddit() string {
	return ldc.Data.Subreddit
}

func (ldc *CommentListingDataChildren) GetScore() float64 {
	return ldc.Data.Score
}

func (ldc *CommentListingDataChildren) GetCreated() float64 {
	return ldc.Data.Created
}

func (ldc *CommentListingDataChildren) GetAuthor() string {
	return ldc.Data.Author
}

func (ldc *CommentListingDataChildren) IsComment() bool {
	return ldc.Kind == "t1"
}

func (ldc *CommentListingDataChildren) IsCommentReply() bool {
	return ldc.Data.Subject == "comment reply"
}
