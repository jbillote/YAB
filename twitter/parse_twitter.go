package twitter

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
func ParseTwitterLink(session *discordgo.Session, message *discordgo.MessageCreate, url string) {
	fxtwitterURL := fmt.Sprintf("https://api.fxtwitter.com/%s", url)

	resp, err := http.Get(fxtwitterURL)
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

	var tweetInfo fxTwitter
	err = json.Unmarshal(body, &tweetInfo)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to unmarshal Tweet information, err=%s", err))
		return
	}

	embed := util.NewEmbed()
	embed.SetTitle("Original Tweet")
	embed.SetURL(tweetInfo.Tweet.URL)
	embed.SetAuthor(tweetInfo.Tweet.Author.URL,
		fmt.Sprintf("%s (@%s)", tweetInfo.Tweet.Author.UserName, tweetInfo.Tweet.Author.ScreenName),
		tweetInfo.Tweet.Author.AvatarURL)
	embed.SetFooter("Twitter", "https://abs.twimg.com/icons/apple-touch-icon-192x192.png")
	embed.SetTimestamp(time.Unix(int64(tweetInfo.Tweet.Timestamp), 0).Format(time.RFC3339))
	embed.SetDescription(tweetInfo.Tweet.Text)
	embed.SetColor(0x3498db)

	var embeds []*discordgo.MessageEmbed
	var videos []string

	for _, u := range tweetInfo.Tweet.Media.Media {
		if u.Type == "photo" {
			if embed.MessageEmbed.Image == nil {
				embed.SetImage(u.URL)
			} else {
				e := util.NewEmbed()
				e.SetURL(tweetInfo.Tweet.URL)
				e.SetImage(u.URL)

				embeds = append(embeds, e.MessageEmbed)
			}
		} else if u.Type == "video" || u.Type == "gif" {
			videos = append(videos, u.URL)
		}
	}
	embeds = append([]*discordgo.MessageEmbed{embed.MessageEmbed}, embeds...)

	session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
		Embeds: embeds,
		AllowedMentions: &discordgo.MessageAllowedMentions{
			RepliedUser: false,
		},
		Reference: message.Reference(),
	})
	for _, v := range videos {
		session.ChannelMessageSend(message.ChannelID, v)
	}
}
