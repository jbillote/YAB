package commands

import (
    "github.com/bwmarrin/discordgo"
    "github.com/jbillote/YAB/util/logger"
)

func gbvsHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
    log := logger.GetLogger("GBVS Frame Data")

    log.Info("GBVS")

    s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
        Type: discordgo.InteractionResponseChannelMessageWithSource,
        Data: &discordgo.InteractionResponseData{
            Content: "GBVS isn't supported yet (・ω<) テヘペロ",
        },
    })
}
