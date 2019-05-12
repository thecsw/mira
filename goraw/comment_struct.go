package goraw

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
	Data   CommentJsonDataThingData
}

type CommentJsonDataThingData struct {
	Author_flair_background_color string // *
	Total_awards_received         int
	Approved_at_utc               string // *
	Distinguished                 string // *
	Mod_reason_by                 string // *
	Banned_by                     string // *
	Author_flair_type             string
	Removal_reason                string // *
	Link_id                       string
	Author_flair_template_id      string // *
	Likes                         bool
	Replies                       string
	User_reports                  []string
	Saved                         bool
	Id                            string
	Banned_at_utc                 string // *
	Mod_reason_title              string // *
	Gilded                        int
	Archived                      bool
	No_follow                     bool
	Author                        string
	Rte_mode                      string
	Can_mod_post                  bool
	Created_utc                   int
	Send_replies                  bool
	Parent_id                     int
	Score                         int
	Author_fullname               string
	Approved_by                   string // *
	Mod_note                      string // *
	All_awardings                 []string
	Subreddit_id                  string
	Body                          string
	Edited                        bool
	Gildings                      string // Probably another struct like Gilding
	Author_flair_css_class        string // *
	Name                          string
	Author_patreon_flair          bool
	Downs                         int
	Author_flair_richtext         []string
	Is_submitter                  bool
	Collapsed_reason              string // *
	Body_html                     string
	Stickied                      bool
	Can_gild                      bool
	Removed                       bool
	Approved                      bool
	Author_flair_text_color       string // *
	Score_hidden                  bool
	Permalink                     string
	Num_reports                   int
	Locked                        bool
	Report_reasons                []string
	Created                       int
	Subreddit                     string
	Author_flair_text             string // *
	Spam                          bool
	Collapsed                     bool
	Subreddit_name_prefixed       string
	Controversiality              int
	Ignore_reports                bool
	Mod_reports                   []string
	Subreddit_type                string
	Ups                           int
}
