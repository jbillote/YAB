package twitter

type vxTwitter struct {
	Timestamp            int      `json:"date_epoch"`
	MediaURLs            []string `json:"mediaURLs"`
	Text                 string   `json:"text"`
	TweetURL             string   `json:"tweetURL"`
	UserName             string   `json:"user_name"`
	ScreenName           string   `json:"user_screen_name"`
	AuthorProfilePicture string   `json:"user_profile_image_url"`
}
