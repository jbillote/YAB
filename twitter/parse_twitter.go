package twitter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/jbillote/YAB/util"
	"github.com/jbillote/YAB/util/logger"
)

var log = logger.GetLogger("TwitterParser")

/*
* Function: ParseTwitterLink
* Get GIFs and videos from Twitter links
*
* Params:
* url: Twitter URL to get content from
 */
func ParseTwitterLink(session *discordgo.Session, channelID string, url string) {
	r, err := regexp.Compile("(\\bx|\\btwitter)\\.com\\/(\\w{1,15}\\/(status|statuses)\\/\\d{2,20})")
	if err != nil {
		log.Error(fmt.Sprintf("Unable to generate regex, err=%s", err))
		return
	}
	match := r.FindStringSubmatch(url)
	if match == nil {
		log.Error(fmt.Sprintf("No Twitter links found, url=%s", url))
		return
	}
	log.Info(match)

	vxtwitterURL := fmt.Sprintf("https://api.vxtwitter.com/%s", match[2])

	resp, err := http.Get(vxtwitterURL)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to get Tweet information, err=%s", err))
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to read Tweet information, err=%s", err))
		return
	}

	var tweetInfo vxTwitter
	err = json.Unmarshal(body, &tweetInfo)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to unmarshal Tweet information, err=%s", err))
		return
	}

	embed := util.NewEmbed()
	embed.SetTitle(fmt.Sprintf("@%s", tweetInfo.ScreenName))
	embed.SetURL(tweetInfo.TweetURL)
	embed.SetAuthor(fmt.Sprintf("https://twitter.com/%s", tweetInfo.ScreenName),
		fmt.Sprintf("%s (@%s)", tweetInfo.UserName, tweetInfo.ScreenName),
		tweetInfo.AuthorProfilePicture)
	embed.SetFooter("Twitter", "http://i.toukat.moe/twitter_logo.png")
	embed.SetTimestamp(time.Unix(int64(tweetInfo.Timestamp), 0).Format(time.RFC3339))

	splitDescription := strings.Split(tweetInfo.Text, " ")
	trimmedDescription := strings.TrimSpace(strings.ReplaceAll(tweetInfo.Text,
		splitDescription[len(splitDescription)-1], ""))
	embed.SetDescription(trimmedDescription)

	var embeds []*discordgo.MessageEmbed
	var videos []string

	for _, u := range tweetInfo.MediaURLs {
		splitUrl := strings.Split(u, ".")
		extension := splitUrl[len(splitUrl)-1]

		if strings.Contains(extension, "jpg") || strings.Contains(extension, "jpeg") ||
			strings.Contains(extension, "png") {

			if embed.MessageEmbed.Image != nil {
				embed.SetImage(u)
			} else {
				e := util.NewEmbed()
				e.SetURL(tweetInfo.TweetURL)
				e.SetImage(u)

				embeds = append(embeds, e.MessageEmbed)
			}
		} else if strings.Contains(extension, "mp4") {
			videos = append(videos, u)
		}
	}
	embeds = append([]*discordgo.MessageEmbed{embed.MessageEmbed}, embeds...)

	session.ChannelMessageSendEmbeds(channelID, embeds)
	for _, v := range videos {
		session.ChannelMessageSend(channelID, v)
	}
}
