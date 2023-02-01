package commands

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
    {
        Name:        "mbtl",
        Description: "Melty Blood: Type Lumina frame data",
    },
}
