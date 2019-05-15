package mira

type Me struct {
	IsEmployee              bool        `json:"is_employee"`
	SeenLayoutSwitch        bool        `json:"seen_layout_switch"`
	HasVisitedNewProfile    bool        `json:"has_visited_new_profile"`
	PrefNoProfanity         bool        `json:"pref_no_profanity"`
	HasExternalAccount      bool        `json:"has_external_account"`
	PrefGeopopular          string      `json:"pref_geopopular"`
	SeenRedesignModal       bool        `json:"seen_redesign_modal"`
	PrefShowTrending        bool        `json:"pref_show_trending"`
	Subreddit               Subreddit_s `json:"subreddit"`
	IsSponsor               bool        `json:"is_sponsor"`
	GoldExpiration          float64     `json:"gold_expiration"`
	HasGoldSubscription     bool        `json:"has_gold_subscription"`
	NumFriends              float64     `json:"num_friends"`
	Features                MeFeatures  `json:"features"`
	HasAndroidSubscription  bool        `json:"has_android_subscription"`
	Verified                bool        `json:"verified"`
	NewModmailExists        bool        `json:"new_modmail_exists"`
	PrefAutoplay            bool        `json:"pref_autoplay"`
	Coins                   float64     `json:"coins"`
	HasPaypalSubscription   bool        `json:"has_paypal_subscription"`
	HasSubscribedToPremium  bool        `json:"has_subscribed_to_premium"`
	Id                      string      `json:"id"`
	HasStripeSubscription   bool        `json:"has_stripe_subscription"`
	SeenPremiumAdblockModal bool        `json:"seen_premium_adblock_modal"`
	CanCreateSubreddit      bool        `json:"can_create_subreddit"`
	Over18                  bool        `json:"over_18"`
	IsGold                  bool        `json:"is_gold"`
	IsMod                   bool        `json:"is_mod"`
	SuspensionExpirationUtc float64     `json:"suspension_expiration_utc"`
	HasVerifiedEmail        bool        `json:"has_verified_email"`
	IsSuspended             bool        `json:"is_suspended"`
	PrefVideoAutoplay       bool        `json:"pref_video_autoplay"`
	InChat                  bool        `json:"in_chat"`
	InRedesignBeta          bool        `json:"in_redesign_beta"`
	IconImg                 string      `json:"icon_img"`
	HasModMail              bool        `json:"has_mod_mail"`
	PrefNightmode           bool        `json:"pref_nightmode"`
	OauthClientId           bool        `json:"oauth_client_id"`
	HideFromRobots          bool        `json:"hide_from_robots"`
	LinkKarma               float64     `json:"link_karma"`
	ForcePasswordReset      bool        `json:"force_password_reset"`
	InboxCount              float64     `json:"inbox_count"`
	PrefTopKarmaSubreddits  bool        `json:"pref_top_karma_subreddits"`
	HasMail                 bool        `json:"has_mail"`
	PrefShowSnoovatar       bool        `json:"pref_show_snoovatar"`
	Name                    string      `json:"name"`
	PrefClickgadget         float64     `json:"pref_clickgadget"`
	Created                 float64     `json:"created"`
	GoldCreddits            float64     `json:"gold_creddits"`
	HasIosSubscription      bool        `json:"has_ios_subscription"`
	PrefShowTwitter         bool        `json:"pref_show_twitter"`
	InBeta                  bool        `json:"in_beta"`
	CommentKarma            float64     `json:"comment_karma"`
	HasSubscribed           bool        `json:"has_subscribed"`
	SeenSubredditChatFtux   bool        `json:"seen_subreddit_chat_ftux"`
}

type MeFeatures struct {
	RichtextPreviews                          bool         `json:"richtext_previews"`
	DoNotTrack                                bool         `json:"do_not_track"`
	ChatSubreddit                             bool         `json:"chat_subreddit"`
	Chat                                      bool         `json:"chat"`
	SeqRandomizeSort                          bool         `json:"seq_randomize_sort"`
	Sequence                                  bool         `json:"sequence"`
	MwebXpromoRevampV2                        MeSubFeature `json:"mweb_xpromo_revamp_v2"`
	MwebXpromoFloat64erstitialCommentsIos     bool         `json:"mweb_xpromo_float64erstitial_comments_ios"`
	ChatReddarReports                         bool         `json:"chat_reddar_reports"`
	ChatRollout                               bool         `json:"chat_rollout"`
	MwebXpromoFloat64erstitialCommentsAndroid bool         `json:"mwev_xpromo_float64erstitial_comments_android"`
	ChatGroutRollout                          bool         `json:"chat_group_rollout"`
	MwebLinkTab                               MeSubFeature `json:"mweb_link_tab"`
	SpezModal                                 bool         `json:"spez_modal"`
	CommunityAwards                           bool         `json:"community_awards"`
	DefaultSrsHoldout                         MeSubFeature `json:"default_srs_holdout"`
	ChatUserSettings                          bool         `json:"chat_user_settings"`
	DualWriteUserPrefs                        bool         `json:"dual_write_user_prefs"`

	MwebXpromoModalListingClickDailyDismissibleAndroid bool `json:"mweb_xpromo_modal_listing_click_daily_dismissible_ios"`
	MwebXpromoModalListingClickDailyDismisssibleIos    bool `json:"mweb_xpromo_modal_listing_click_daily_dismissible_android"`
}

type MeSubFeature struct {
	Owner        string  `json:"owner"`
	Variant      string  `json:"variant"`
	ExperimentId float64 `json:"experiment_id"`
}
