package utils

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"monbotdiscord/bot/config"
)

type ButtonId string

const (
	BtnRules ButtonId = "rules_accept"
)

func HandleBtnExecute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.MessageComponentData().CustomID == string(BtnRules) {
		err := session.GuildMemberRoleAdd(interaction.GuildID, interaction.Member.User.ID, config.ConfigBotInstance.Guild.Roles.Validated)
		if err != nil {
			log.Println("Erreur lors de l'ajout du rôle", err)
			return
		}
		rulesResponse := discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Merci d'avoir accepté les règles",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		}

		err = session.InteractionRespond(interaction.Interaction, &rulesResponse)

	}
}
