package mira

type CommentWrap struct {
	Json CommentJson `json:"json"`
}

type CommentJson struct {
	Errors []string        `json:"errors"`
	Data   CommentJsonData `json:"data"`
}

type CommentJsonData struct {
	Things []CommentJsonDataThing `json:"things"`
}

type CommentJsonDataThing struct {
	Kind string                   `json:"kind"`
	Data CommentJsonDataThingData `json:"data"`
}

type CommentJsonDataThingData struct {
	AuthorFlairBackgroundColor string   `json:"author_flair_background_color"`
	TotalAwardsReceived        float64  `json:"total_awards_received"`
	ApprovedAtUtc              string   `json:"approved_at_utc"`
	Distinguished              string   `json:"distinguished"`
	ModReasonBy                string   `json:"mod_reason_by"`
	BannedBy                   string   `json:"banned_by"`
	AuthorFlairType            string   `json:"author_flair_type"`
	RemovalReason              string   `json:"removal_reason"`
	LinkId                     string   `json:"link_id"`
	AuthorFlairTemplateId      string   `json:"author_flair_template_id"`
	Likes                      bool     `json:"likes"`
	Replies                    string   `json:"replies"`
	UserReports                []string `json:"user_reports"`
	Saved                      bool     `json:"saved"`
	Id                         string   `json:"id"`
	BannedAtUtc                string   `json:"banned_at_utc"`
	ModReasonTitle             string   `json:"mod_reason_title"`
	Gilded                     float64  `json:"gilded"`
	Archived                   bool     `json:"archived"`
	NoFollow                   bool     `json:"no_follow"`
	Author                     string   `json:"author"`
	RteMode                    string   `json:"rte_mode"`
	CanModPost                 bool     `json:"can_mod_post"`
	CreatedUtc                 float64  `json:"created_utc"`
	SendReplies                bool     `json:"send_replies"`
	ParentId                   string   `json:"parent_id"`
	Score                      float64  `json:"score"`
	AuthorFullname             string   `json:"author_fullname"`
	ApprovedBy                 string   `json:"approved_by"`
	Mod_note                   string   `json:"mod_note"`
	AllAwardings               []string `json:"all_awardings"`
	SubredditId                string   `json:"subreddit_id"`
	Body                       string   `json:"body"`
	Edited                     bool     `json:"edited"`
	Gildings                   Gilding  `json:"gildings"`
	AuthorFlairCssClass        string   `json:"author_flair_css_class"`
	Name                       string   `json:"name"`
	AuthorPatreonFlair         bool     `json:"author_patreon_flair"`
	Downs                      float64  `json:"downs"`
	AuthorFlairRichtext        []string `json:"author_flair_richtext"`
	IsSubmitter                bool     `json:"is_submitter"`
	CollapsedReason            string   `json:"collapsed_reason"`
	BodyHtml                   string   `json:"body_html"`
	Stickied                   bool     `json:"stickied"`
	CanGild                    bool     `json:"can_gild"`
	Removed                    bool     `json:"removed"`
	Approved                   bool     `json:"approved"`
	AuthorFlairTextColor       string   `json:"author_flair_text_color"`
	ScoreHidden                bool     `json:"score_hidden"`
	Permalink                  string   `json:"permalink"`
	NumReports                 float64  `json:"num_reports"`
	Locked                     bool     `json:"locked"`
	ReportReasons              []string `json:"report_reasons"`
	Created                    float64  `json:"created"`
	Subreddit                  string   `json:"subreddit"`
	AuthorFlairText            string   `json:"author_flair_text"`
	Spam                       bool     `json:"spam"`
	Collapsed                  bool     `json:"collapsed"`
	SubredditNamePrefixed      string   `json:"subreddit_name_prefixed"`
	Controversiality           float64  `json:"controversiality"`
	IgnoreReports              bool     `json:"ignore_reports"`
	ModReports                 []string `json:"mod_reports"`
	SubredditType              string   `json:"subreddit_type"`
	Ups                        float64  `json:"ups"`
}

type Gilding struct {
	Gid map[string]int `json:"gid"`
}
