package mira

type Comment struct {
	Json CommentJson
}

type CommentJson struct {
	Errors []string
	Data   CommentJsonData
}

type CommentJsonData struct {
	Things []CommentJsonDataThing
}

type CommentJsonDataThing struct {
	Kind string
	Data CommentJsonDataThingData
}

type CommentJsonDataThingData struct {
	AuthorFlairBackgroundColor string // *
	TotalAwardsReceived        float64
	ApprovedAtUtc              string // *
	Distinguished              string // *
	ModReasonBy                string // *
	BannedBy                   string // *
	AuthorFlairType            string
	RemovalReason              string // *
	LinkId                     string
	AuthorFlairTemplateId      string // *
	Likes                      bool
	Replies                    string
	UserReports                []string
	Saved                      bool
	Id                         string
	BannedAtUtc                string // *
	ModReasonTitle             string // *
	Gilded                     float64
	Archived                   bool
	NoFollow                   bool
	Author                     string
	RteMode                    string
	CanModPost                 bool
	CreatedUtc                 float64
	SendReplies                bool
	ParentId                   float64
	Score                      float64
	AuthorFullname             string
	ApprovedBy                 string // *
	Mod_note                   string // *
	AllAwardings               []string
	SubredditId                string
	Body                       string
	Edited                     bool
	Gildings                   string // Probably another struct like Gilding
	AuthorFlairCssClass        string // *
	Name                       string
	AuthorPatreonFlair         bool
	Downs                      float64
	AuthorFlairRichtext        []string
	IsSubmitter                bool
	CollapsedReason            string // *
	BodyHtml                   string
	Stickied                   bool
	CanGild                    bool
	Removed                    bool
	Approved                   bool
	AuthorFlairTextColor       string // *
	ScoreHidden                bool
	Permalink                  string
	NumReports                 float64
	Locked                     bool
	ReportReasons              []string
	Created                    float64
	Subreddit                  string
	AuthorFlairText            string // *
	Spam                       bool
	Collapsed                  bool
	SubredditNamePrefixed      string
	Controversiality           float64
	IgnoreReports              bool
	ModReports                 []string
	SubredditType              string
	Ups                        float64
}
