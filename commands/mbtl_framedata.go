package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jbillote/YAB/graphql"
	"github.com/jbillote/YAB/util"
	"github.com/jbillote/YAB/util/logger"
)

func mbtlHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log := logger.GetLogger("MBTL Frame Data")

	log.Info("Fetching options")
	options := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(i.ApplicationCommandData().Options))
	for _, o := range i.ApplicationCommandData().Options {
		options[o.Name] = o
	}

	move, _ := graphql.QueryMBTLMove(options["character"].StringValue(), options["input"].StringValue())

	embed := util.NewEmbed()
	embed.SetTitle(move.Name)
	embed.AddField("Input", move.Input, true)
	embed.AddBlankField(true)
	embed.AddBlankField(true)
	embed.AddField("Damage", move.Damage, true)
	embed.AddField("Block", move.Block, true)
	embed.AddField("Cancel", move.Cancel, true)
	embed.AddField("Property", move.Property, true)
	embed.AddField("Cost", move.Cost, true)
	embed.AddField("Attribute", move.Attribute, true)
	embed.AddField("Startup", move.Startup, true)
	embed.AddField("Active", move.Active, true)
	embed.AddField("Recovery", move.Recovery, true)
	embed.AddField("Overall", move.Overall, true)
	embed.AddField("Advantage", move.Advantage, true)
	embed.AddField("Invuln", move.Invuln, true)

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				embed.MessageEmbed,
			},
		},
	})
}
