package models

func (l *CommentListing) GetChildren() []Comment { return l.Data.Children }
func (ldc Comment) GetId() string                { return ldc.Data.Name }
func (ldc Comment) GetParentId() string          { return ldc.Data.ParentId }
func (ldc Comment) IsRoot() bool                 { return string(ldc.Data.ParentId[1]) == "3" }
func (ldc Comment) GetTitle() string             { return ldc.Data.Body }
func (ldc Comment) GetBody() string              { return ldc.Data.Body }
func (ldc Comment) GetSubreddit() string         { return ldc.Data.Subreddit }
func (ldc Comment) GetUps() float64              { return ldc.Data.Score }
func (ldc Comment) GetDowns() float64            { return 0 }
func (ldc Comment) GetCreated() float64          { return ldc.Data.Created }
func (ldc Comment) GetAuthor() string            { return ldc.Data.Author }
func (ldc Comment) IsComment() bool              { return ldc.Kind == "t1" }
func (ldc Comment) IsCommentReply() bool         { return ldc.Data.Subject == "comment reply" }
func (ldc Comment) IsMention() bool              { return ldc.Data.Subject == "username mention" }
func (ldc Comment) GetName() string              { return ldc.Data.Name }
func (ldc Comment) GetKarma() float64            { return ldc.Data.Score }
func (ldc Comment) GetUrl() string               { return ldc.Data.Permalink }
func (ldc Comment) GetFlair() string             { return ldc.Data.Context }
