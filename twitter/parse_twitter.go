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
	// Give time for embeds to load
	time.Sleep(5 * time.Second)

	// Make sure embeds are up to date
	m, err := session.ChannelMessage(message.ChannelID, message.ID)
	if err != nil {
		log.Error(err)
		return
	}

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

	if len(m.Embeds) < 1 || (tweetInfo.Tweet.Media != nil &&
		(len(tweetInfo.Tweet.Media.Photos) > 1 ||
			tweetInfo.Tweet.Media.Videos != nil)) {

		originalEmbeds, originalVideos := generateTweetEmbeds(tweetInfo.Tweet, false)
		var quoteEmbeds []*discordgo.MessageEmbed
		var quoteVideos []string

		if tweetInfo.Tweet.Quote != nil {
			quoteEmbeds, quoteVideos = generateTweetEmbeds(*tweetInfo.Tweet.Quote, true)
		}

		session.ChannelMessageSendComplex(message.ChannelID, &discordgo.MessageSend{
			Embeds: originalEmbeds,
			AllowedMentions: &discordgo.MessageAllowedMentions{
				RepliedUser: false,
			},
			Reference: message.Reference(),
		})
		for _, v := range originalVideos {
			session.ChannelMessageSend(message.ChannelID, v)
		}

		session.ChannelMessageSendEmbeds(message.ChannelID, quoteEmbeds)
		for _, v := range quoteVideos {
			session.ChannelMessageSend(message.ChannelID, v)
		}
	}
}

func generateTweetEmbeds(tweetInfo tweet, isQuote bool) ([]*discordgo.MessageEmbed, []string) {
	var embeds []*discordgo.MessageEmbed
	var videos []string

	embed := util.NewEmbed()
	if isQuote {
		embed.SetTitle("Quoted Tweet")
	} else {
		embed.SetTitle("Original Tweet")
	}
	embed.SetURL(tweetInfo.URL)
	embed.SetAuthor(tweetInfo.Author.URL,
		fmt.Sprintf("%s (@%s)", tweetInfo.Author.UserName, tweetInfo.Author.ScreenName),
		tweetInfo.Author.AvatarURL)
	embed.SetFooter("Twitter", "https://abs.twimg.com/icons/apple-touch-icon-192x192.png")
	embed.SetTimestamp(time.Unix(int64(tweetInfo.Timestamp), 0).Format(time.RFC3339))
	embed.SetDescription(tweetInfo.Text)
	embed.SetColor(0x3498db)

	for _, u := range tweetInfo.Media.Media {
		if u.Type == "photo" {
			if embed.MessageEmbed.Image == nil {
				embed.SetImage(u.URL)
			} else {
				e := util.NewEmbed()
				e.SetURL(tweetInfo.URL)
				e.SetImage(u.URL)

				embeds = append(embeds, e.MessageEmbed)
			}
		} else if u.Type == "video" || u.Type == "gif" {
			videos = append(videos, u.URL)
		}
	}
	embeds = append([]*discordgo.MessageEmbed{embed.MessageEmbed}, embeds...)

	return embeds, videos
}
