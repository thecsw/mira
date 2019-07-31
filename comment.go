package mira

func (c CommentWrap) GetId() string          { return c.Json.Data.Things[0].Data.Name }
func (c CommentWrap) GetSubredditId() string { return c.Json.Data.Things[0].Data.SubredditId }
func (c CommentWrap) GetParentId() string    { return c.Json.Data.Things[0].Data.ParentId }
func (c CommentWrap) GetAuthor() string      { return c.Json.Data.Things[0].Data.Author }
func (c CommentWrap) GetAuthorId() string    { return c.Json.Data.Things[0].Data.AuthorFullname }
func (c CommentWrap) GetSubreddit() string   { return c.Json.Data.Things[0].Data.Subreddit }
func (c CommentWrap) CreatedAt() float64     { return c.Json.Data.Things[0].Data.Created }
func (c CommentWrap) GetBody() string        { return c.Json.Data.Things[0].Data.Body }
func (c CommentWrap) GetScore() float64      { return c.Json.Data.Things[0].Data.Score }
func (c CommentWrap) GetUps() float64        { return c.Json.Data.Things[0].Data.Ups }
func (c CommentWrap) GetDowns() float64      { return c.Json.Data.Things[0].Data.Downs }
func (c CommentWrap) IsSticky() bool         { return c.Json.Data.Things[0].Data.Stickied }
func (c CommentWrap) IsRemoved() bool        { return c.Json.Data.Things[0].Data.Removed }
func (c CommentWrap) IsApproved() bool       { return c.Json.Data.Things[0].Data.Approved }
func (c CommentWrap) IsAuthor() bool         { return c.Json.Data.Things[0].Data.IsSubmitter }
