package commands

import (
    "fmt"
    "github.com/bwmarrin/discordgo"
    "github.com/jbillote/YAB/util/logger"
)

func mbtlHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    log := logger.GetLogger("MBTL Frame Data")

    log.Info("Fetching options")
    options := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(i.ApplicationCommandData().Options))
    for _, o := range i.ApplicationCommandData().Options {
        options[o.Name] = o
    }

    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: fmt.Sprintf("%v's %v", options["character"].StringValue(), options["input"].StringValue()),
        },
    })
}
