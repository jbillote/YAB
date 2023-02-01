package commands

import "github.com/bwmarrin/discordgo"

var Commands = []*discordgo.ApplicationCommand{
    {
        Name:        "mbtl",
        Description: "Melty Blood: Type Lumina frame data",
        Options: []*discordgo.ApplicationCommandOption{
            {
                Type:        discordgo.ApplicationCommandOptionString,
                Name:        "character",
                Description: "Character Name",
                Required:    true,
            },
            {
                Type:        discordgo.ApplicationCommandOptionString,
                Name:        "input",
                Description: "Move input",
                Required:    true,
            },
        },
    },
    {
        Name:        "gbvs",
        Description: "Granblue Fantasy: Versus frame data",
    },
}
