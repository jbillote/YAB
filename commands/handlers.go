package commands

import "github.com/bwmarrin/discordgo"

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
    "mbtl": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
        s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
            Type: discordgo.InteractionResponseChannelMessageWithSource,
            Data: &discordgo.InteractionResponseData{
                Content: "a",
            },
        })
    },
}
