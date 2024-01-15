package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"monbotdiscord/bot/utils"
	"strconv"
)

var defaultPermissions int64 = discordgo.PermissionManageMessages

var CmdClear = discordgo.ApplicationCommand{
	Name:                     "clear",
	Description:              "Supprime un nombre de messages donné",
	DefaultMemberPermissions: &defaultPermissions,
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "nombre",
			Description: "Nombre de messages à supprimer",
			Type:        discordgo.ApplicationCommandOptionInteger,
			Required:    true,
		},
	},
}

func CmdClearResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) discordgo.InteractionResponse {

	// Get the number of messages to delete
	number := interaction.ApplicationCommandData().Options[0].IntValue()
	if number > 100 {
		number = 100
	}

	channel, err := session.Channel(interaction.ChannelID)
	if err != nil {
		log.Println("Erreur lors de la récupération du channel", err)
	}

	utils.ClearChannelWithCount(session, channel, int(number))

	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,

		Data: &discordgo.InteractionResponseData{
			Content: "Vous avez clear " + strconv.FormatInt(number, 10) + " messages",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	}

	return response
}
