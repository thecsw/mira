package mira

func (r *Redditor) IsEmployee() bool       { return r.Data.IsEmployee }
func (r *Redditor) GetName() string        { return r.Data.Name }
func (r *Redditor) GetAuthor() string      { return r.Data.Name }
func (r *Redditor) GetId() string          { return r.Kind + "_" + r.Data.Id }
func (r *Redditor) GetParentId() string    { return r.Kind + "_" + r.Data.Id }
func (r *Redditor) GetDescription() string { return r.Data.Subreddit.PublicDescription }
func (r *Redditor) GetCreated() float64    { return r.Data.Created }
func (r *Redditor) GetKarma() float64      { return r.Data.LinkKarma + r.Data.CommentKarma }
func (r *Redditor) GetUps() float64        { return r.Data.LinkKarma }
func (r *Redditor) GetDowns() float64      { return r.Data.CommentKarma }
func (r *Redditor) GetBody() string        { return r.Data.Subreddit.PublicDescription }
func (r *Redditor) GetTitle() string       { return r.Data.Subreddit.Title }
func (r *Redditor) GetFlair() string       { return r.Data.Subreddit.PublicDescription }
func (r *Redditor) GetSubreddit() string   { return r.Data.Subreddit.Name }
func (r *Redditor) GetUrl() string         { return r.Data.Subreddit.IconImg }
func (r *Redditor) IsRoot() bool           { return true }
