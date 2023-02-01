package main

import (
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
        if util.URLValid(m) && util.URLAvailable(m) {
            twitter.ParseTwitterLink(session, message.ChannelID, m)
        }
    }
}
