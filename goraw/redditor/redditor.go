package redditor

type Redditor struct {
	Kind string `json:"kind"`
	Data Data_s `json:"data"`
}

type Data_s struct {
	IsEmployee        bool        `json:"is_employee"`
	IconImg           string      `json:"icon_img"`
	PrefShowSnoovatar bool        `json:"pref_show_snoovatar"`
	Name              string      `json:"name"`
	IsFriend          bool        `json:"is_friend"`
	Created           float64     `json:"created"`
	HasSubscribed     bool        `json:"has_subscribed"`
	HideFromRobots    bool        `json:"hide_from_robots"`
	CreatedUtc        float64     `json:"created_utc"`
	LinkKarma         float64     `json:"link_karma"`
	CommentKarma      float64     `json:"comment_karma"`
	IsGold            bool        `json:"is_gold"`
	IsMod             bool        `json:"is_mod"`
	Verified          bool        `json:"verified"`
	Subreddit         Subreddit_s `json:"subreddit"`
	HasVerifiedEmail  bool        `json:"has_verified_email"`
	Id                string      `json:"id"`
}

type Me struct {
	IsEmployee                bool        `json:"is_employee"`
	SeenLayoutSwitch          bool        `json:"seen_layout_switch"`
	HasVisitedNewProfile      bool        `json:"has_visited_new_profile"`
	PrefNoProfanity           bool        `json:"pref_no_profanity"`
	HasExternalAccount        bool        `json:"has_external_account"`
	PrefGeopopular            string      `json:"pref_geopopular"`
	SeenRedesignModal         bool        `json:"seen_redesign_modal"`
	PrefShowTrending          bool        `json:"pref_show_trending"`
	Subreddit                 Subreddit_s `json:"subreddit"`
	IsSponsor                 bool        `json:"is_sponsor"`
	GoldExpiration            int64       `json:"gold_expiration"`
	HasGoldSubscription       bool        `json:"has_gold_subscription`
	NumFriends                int64       `json:"num_friends"`
	Features                  Features_s  `json:"features"`
	HasAndroidSubscription    bool        `json:"has_android_subscription"`
	Verified                  bool        `json:"verified"`
	NewModmailExists          bool        `json:"new_modmail_exists"`
	PrefAutoplay              bool        `json:"pref_autoplay"`
	Coins                     int64       `json:"coins"`
	HasPaypalSubscription     bool        `json:"has_paypal_subscription"`
	HasSubscribedToPremium    bool        `json:"has_subscribed_to_premium"`
	Id                        string      `json:"id"`
	HasStripeSubscription     bool        `json:"has_stripe_subscription"`
	SeenPremiumAdblockModal   bool        `json:"seen_premium_adblock_modal"`
	CanCreateSubreddit        bool        `json:"can_create_subreddit"`
	Over18                    bool        `json:"over_18"`
	IsGold                    bool        `json:"is_gold"`
	IsMod                     bool        `json:"is_mod"`
	SuspensionExpirationUtc   int64       `json:"suspension_expiration_utc"`
	HasVerifiedEmail          bool        `json:"has_verified_email"`
	IsSuspended               bool        `json:"is_suspended"`
	PrefVideoAutoplay         bool        `json:"pref_video_autoplay"`
	InChat                    bool        `json:"in_chat"`
	InRedesignBeta            bool        `json:"in_redesign_beta"`
	IconImg                   string      `json:"icon_img"`
	HasModMail                bool        `json:"has_mod_mail"`
	PrefNightmode             bool        `json:"pref_nightmode"`
	OauthClientId             bool        `json:"oauth_client_id"`
	HideFromRobots            bool        `json:"hide_from_robots"`
	LinkKarma                 int64       `json:"link_karma"`
	ForcePasswordReset        bool        `json:"force_password_reset"`
	InboxCount                int64       `json:"inbox_count"`
	PrefTopKarmaSubreddits    bool        `json:"pref_top_karma_subreddits"`
	HasMail                   bool        `json:"has_mail"`
	PrefShowSnoovatar         bool        `json:"pref_show_snoovatar"`
	Name                      string      `json:"name"`
	PrefClickgadget           int64       `json:"pref_clickgadget"`
	Created                   int64       `json:"created"`
	GoldCreddits              int64       `json:"gold_creddits"`
	HasIosSubscription        bool        `json:"has_ios_subscription"`
	PrefShowTwitter           bool        `json:"pref_show_twitter"`
	InBeta                    bool        `json:"in_beta"`
	CommentKarma              int64       `json:"comment_karma"`
	HasSubscribed             bool        `json:"has_subscribed"`
	SeenSubredditChatFtux     bool        `json:"seen_subreddit_chat_ftux"`
}

type Features_s struct {
	RichtextPreviews                    bool         `json:"richtext_previews"`
	DoNotTrack                          bool         `json:"do_not_track"`
	ChatSubreddit                       bool         `json:"chat_subreddit"`
	Chat                                bool         `json:"chat"`
	SeqRandomizeSort                    bool         `json:"seq_randomize_sort"`
	Sequence                            bool         `json:"sequence"`
	MwebXpromoRevampV2                 	SubFeature_s `json:"mweb_xpromo_revamp_v2"`
	MwebXpromoInterstitialCommentsIos	bool         `json:"mweb_xpromo_interstitial_comments_ios"`
	ChatReddarReports                   bool         `json:chat_reddar_reports`
	ChatRollout                         bool         `json:"chat_rollout"`
	MwebXpromoInterstitialCommentsAndroid bool       `json:"mwev_xpromo_interstitial_comments_android"`
	ChatGroutRollout                    bool         `json:"chat_group_rollout"`
	MwebLinkTab                         SubFeature_s `json:"mweb_link_tab"`
	SpezModal                           bool         `json:"spez_modal"`
	CommunityAwards                     bool         `json:"community_awards"`
	DefaultSrsHoldout                   SubFeature_s `json:"default_srs_holdout"`
	ChatUserSettings                    bool         `json:"chat_user_settings"`
	DualWriteUserPrefs                  bool         `json:"dual_write_user_prefs"`

	MwebXpromoModalListingClickDailyDismissibleAndroid    bool    `json:"mweb_xpromo_modal_listing_click_daily_dismissible_ios"`
	MwebXpromoModalListingClickDailyDismisssibleIos       bool    `json:"mweb_xpromo_modal_listing_click_daily_dismissible_android"` 
}

type SubFeature_s struct {
	Owner                 string      `json:"owner"`
	Variant               string      `json:"variant"`
	ExperimentId          int64       `json:"experiment_id"`
}

type Subreddit_s struct {
	DefaultSet                 bool      `json:"default_set"`
	UserIsContributor          bool      `json:"user_is_contributor"`
	BannerImg                  string    `json:"banner_img"`
	DisableContributorRequests bool      `json:"disable_contributor_requests"`
	UserIsBanned               bool      `json:"user_is_banned"`
	FreeFormReports            bool      `json:"free_form_reports"`
	CommunityIcon              string    `json:"community_icon"`
	ShowMedia                  bool      `json:"show_media"`
	IconColor                  string    `json:"icon_color"`
	UserIsMuted                bool      `json:"user_is_muted"`
	DisplayName                string    `json:"display_name"`
	HeaderImg                  *string   `json:"header_img"`
	Title                      string    `json:"title"`
	Over18                     bool      `json:"over_18"`
	IconSize                   []float64 `json:"icon_size"`
	PrimaryColor               string    `json:"primary_color"`
	IconImg                    string    `json:"icon_img"`
	Description                string    `json:"description"`
	HeaderSize                 *string   `json:"header_size"`
	RestrictPosting            bool      `json:"restrict_posting"`
	RestrictCommenting         bool      `json:"restrict_commenting"`
	Subscribers                float64   `json:"subscribers"`
	IsDefaultIcon              bool      `json:"is_default_icon"`
	LinkFlairPosition          string    `json:"link_flair_position"`
	DisplayNamePrefixed        string    `json:"display_name_prefixed"`
	KeyColor                   string    `json:"key_color"`
	Name                       string    `json:"name"`
	IsDefaultBanner            bool      `json:"is_default_banner"`
	Url                        string    `json:"url"`
	BannerSize                 []float64 `json:"banner_size"`
	UserIsModerator            bool      `json:"user_is_moderator"`
	PublicDescription          string    `json:"public_description"`
	LinkFlairEnabled           bool      `json:"link_flair_enabled"`
	SubredditType              string    `json:"subreddit_type"`
	UserIsSubscriber           bool      `json:"user_is_subscriber"`
}
