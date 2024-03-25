package twitter

type fxTwitter struct {
	Tweet tweet `json:"tweet"`
}

type tweet struct {
	URL       string `json:"url"`
	Text      string `json:"text"`
	Author    author `json:"author"`
	Timestamp int    `json:"created_timestamp"`
	Media     media  `json:"media"`
	Quote     *tweet `json:"quote"`
}

type author struct {
	UserName   string `json:"name"`
	ScreenName string `json:"screen_name"`
	AvatarURL  string `json:"avatar_url"`
	URL        string `json:"url"`
}

type media struct {
	Media  []attachment `json:"all"`
	Videos []attachment `json:"videos"`
	Photos []attachment `json:"photos"`
}

type attachment struct {
	URL  string `json:"url"`
	Type string `json:"type"`
}
