package commands

import "github.com/bwmarrin/discordgo"

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
    "mbtl": mbtlHandler,
}
