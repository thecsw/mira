package models

func (c CommentWrap) getThing() CommentJsonDataThing {
	if len(c.Json.Data.Things) > 0 {
		return c.Json.Data.Things[0]
	}
	return CommentJsonDataThing{}
}
func (c CommentWrap) GetId() string          { return c.getThing().Data.Name }
func (c CommentWrap) GetSubredditId() string { return c.getThing().Data.SubredditId }
func (c CommentWrap) GetParentId() string    { return c.getThing().Data.ParentId }
func (c CommentWrap) GetAuthor() string      { return c.getThing().Data.Author }
func (c CommentWrap) GetAuthorId() string    { return c.getThing().Data.AuthorFullname }
func (c CommentWrap) GetSubreddit() string   { return c.getThing().Data.Subreddit }
func (c CommentWrap) CreatedAt() float64     { return c.getThing().Data.Created }
func (c CommentWrap) GetBody() string        { return c.getThing().Data.Body }
func (c CommentWrap) GetScore() float64      { return c.getThing().Data.Score }
func (c CommentWrap) GetUps() float64        { return c.getThing().Data.Ups }
func (c CommentWrap) GetDowns() float64      { return c.getThing().Data.Downs }
func (c CommentWrap) IsSticky() bool         { return c.getThing().Data.Stickied }
func (c CommentWrap) IsRemoved() bool        { return c.getThing().Data.Removed }
func (c CommentWrap) IsApproved() bool       { return c.getThing().Data.Approved }
func (c CommentWrap) IsAuthor() bool         { return c.getThing().Data.IsSubmitter }
