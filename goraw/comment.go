package goraw

// Get ID of the comment. Should be something "t1_..."
func (c Comment) GetId() string {
	return c.Json.Data.Things[0].Data.Name
}

// Get the name of the author. With no u/ preppended
func (c Comment) GetAuthor() string {
	return c.Json.Data.Things[0].Data.Author
}

// Get the subreddit's name. With no r/ preppended
func (c Comment) GetSubreddit() string {
	return c.Json.Data.Things[0].Data.Subreddit
}

// Get the UNIX timestamp when the comment was created
func (c Comment) CreatedAt() int {
	return c.Json.Data.Things[0].Data.Created
}

// Get the body of the message
func (c Comment) GetBody() string {
	return c.Json.Data.Things[0].Data.Body
}

// Get the score of the comment (Ups - Downs)
func (c Comment) GetScore() int {
	return c.Json.Data.Things[0].Data.Score
}

// Get the number of upvotes on the comment
func (c Comment) GetUps() int {
	return c.Json.Data.Things[0].Data.Ups
}

// Get the number of downvotes
func (c Comment) GetDowns() int {
	return c.Json.Data.Things[0].Data.Downs
}

// Return true if the comment is stickied, false otherwise
func (c Comment) IsSticky() bool {
	return c.Json.Data.Things[0].Data.Stickied
}

// Return true if the comment is removed, false otherwise
func (c Comment) IsRemoved() bool {
	return c.Json.Data.Things[0].Data.Removed
}

// Return true if the comment is approved, false otherwise
func (c Comment) IsApproved() bool {
	return c.Json.Data.Things[0].Data.Approved
}

// Return true if the author is submission's author, false otherwise
func (c Comment) IsAuthor() bool {
	return c.Json.Data.Things[0].Data.IsSubmitter
}
