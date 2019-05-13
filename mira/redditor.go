package mira

//
func (r *Redditor) IsEmployee() bool {
	return r.Data.IsEmployee
}

//
func (r *Redditor) GetName() string {
	return r.Data.Name
}

//
func (r *Redditor) GetId() string {
	return r.Kind + r.Data.Id
}

//
func (r *Redditor) GetDescription() string {
	return r.Data.Subreddit.PublicDescription
}

//
func (r *Redditor) GetCreated() float64 {
	return r.Data.Created
}

//
func (r *Redditor) GetLinkKarma() float64 {
	return r.Data.LinkKarma
}

//
func (r *Redditor) GetCommentKarma() float64 {
	return r.Data.CommentKarma
}
