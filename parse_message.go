package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/jbillote/YAB/twitter"
	"github.com/jbillote/YAB/util"
)

func ParseMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == uid {
		return
	}

	splitMessage := strings.Split(message.Content, " ")

	// Check for Twitter links
	for _, m := range splitMessage {
		if util.URLValid(m) {
			r, err := regexp.Compile(`(\bx|\btwitter)\.com\/(\w{1,15}\/(status|statuses)\/\d{2,20})`)
			if err != nil {
				log.Error(fmt.Sprintf("Unable to generate regex, err=%s", err))
				return
			}
			match := r.FindStringSubmatch(m)
			if match == nil {
				log.Error(fmt.Sprintf("No Twitter links found, url=%s", m))
				return
			}
			log.Info(match)
			log.Info(match[2])

			twitter.ParseTwitterLink(session, message, match[2])
		}
	}
}
