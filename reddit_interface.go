package mira

type MiraInterface interface {
	GetId() string
	GetParentId() string
	GetTitle() string
	GetBody() string
	GetAuthor() string
	GetName() string
	GetKarma() float64
	GetUps() float64
	GetDowns() float64
	GetSubreddit() string
	GetCreated() float64
	GetFlair() string
	GetUrl() string
	IsRoot() bool
}
