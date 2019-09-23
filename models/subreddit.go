package models

func (s Subreddit) GetId() string                { return s.Data.Name }
func (s Subreddit) GetParentId() string          { return s.Data.Name }
func (s Subreddit) GetName() string              { return s.Data.Title }
func (s Subreddit) GetAuthor() string            { return s.Data.Title }
func (s Subreddit) GetSubreddit() string         { return s.Data.Name }
func (s Subreddit) GetTitle() string             { return s.Data.Title }
func (s Subreddit) GetBody() string              { return s.Data.Description }
func (s Subreddit) GetDisplayName() string       { return s.Data.DisplayName }
func (s Subreddit) GetUrl() string               { return s.Data.Url }
func (s Subreddit) GetUps() float64              { return 0 }
func (s Subreddit) GetKarma() float64            { return 0 }
func (s Subreddit) GetDowns() float64            { return 0 }
func (s Subreddit) IsOver18() bool               { return s.Data.Over18 }
func (s Subreddit) GetPublicDescription() string { return s.Data.PublicDescription }
func (s Subreddit) GetDescription() string       { return s.Data.Description }
func (s Subreddit) GetFlair() string             { return s.Data.HeaderTitle }
func (s Subreddit) GetCreated() float64          { return s.Data.CreatedUtc }
func (s Subreddit) GetSubscribers() float64      { return s.Data.Subscribers }
func (s Subreddit) IsRoot() bool                 { return true }
