package mira

// Get ID of the comment. Should be something "t1_..."
func (c CommentWrap) GetId() string {
	return c.Json.Data.Things[0].Data.Name
}

func (c CommentWrap) GetSubredditId() string {
	return c.Json.Data.Things[0].Data.SubredditId
}

// Get ID of the comment. Should be something "t1_..."
func (c CommentWrap) GetParentId() string {
	return c.Json.Data.Things[0].Data.ParentId
}

// Get the name of the author. With no u/ preppended
func (c CommentWrap) GetAuthor() string {
	return c.Json.Data.Things[0].Data.Author
}

// Get the name of the author. With no u/ preppended
func (c CommentWrap) GetAuthorId() string {
	return c.Json.Data.Things[0].Data.AuthorFullname
}

// Get the subreddit's name. With no r/ preppended
func (c CommentWrap) GetSubreddit() string {
	return c.Json.Data.Things[0].Data.Subreddit
}

// Get the UNIX timestamp when the comment was created
func (c CommentWrap) CreatedAt() float64 {
	return c.Json.Data.Things[0].Data.Created
}

// Get the body of the message
func (c CommentWrap) GetBody() string {
	return c.Json.Data.Things[0].Data.Body
}

// Get the score of the comment (Ups - Downs)
func (c CommentWrap) GetScore() float64 {
	return c.Json.Data.Things[0].Data.Score
}

// Get the number of upvotes on the comment
func (c CommentWrap) GetUps() float64 {
	return c.Json.Data.Things[0].Data.Ups
}

// Get the number of downvotes
func (c CommentWrap) GetDowns() float64 {
	return c.Json.Data.Things[0].Data.Downs
}

// Return true if the comment is stickied, false otherwise
func (c CommentWrap) IsSticky() bool {
	return c.Json.Data.Things[0].Data.Stickied
}

// Return true if the comment is removed, false otherwise
func (c CommentWrap) IsRemoved() bool {
	return c.Json.Data.Things[0].Data.Removed
}

// Return true if the comment is approved, false otherwise
func (c CommentWrap) IsApproved() bool {
	return c.Json.Data.Things[0].Data.Approved
}

// Return true if the author is submission's author, false otherwise
func (c CommentWrap) IsAuthor() bool {
	return c.Json.Data.Things[0].Data.IsSubmitter
}
