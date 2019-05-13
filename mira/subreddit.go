package mira

func (s Subreddit) GetId() string {
	return s.Data.Name
}

func (s Subreddit) GetName() string {
	return s.Data.Title
}

func (s Subreddit) GetDisplayName() string {
	return s.Data.DisplayName
}

func (s Subreddit) GetUrl() string {
	return s.Data.Url
}

func (s Subreddit) IsOver18() bool {
	return s.Data.Over18
}

func (s Subreddit) GetPublicDescription() string {
	return s.Data.PublicDescription
}

func (s Subreddit) GetDescription() string {
	return s.Data.Description
}

func (s Subreddit) GetCreated() float64 {
	return s.Data.CreatedUtc
}

func (s Subreddit) GetSubscribers() float64 {
	return s.Data.Subscribers
}
