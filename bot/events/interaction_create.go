package events

import (
	"github.com/bwmarrin/discordgo"
	"monbotdiscord/bot/commands"
	"monbotdiscord/bot/utils"
)

var InteractionCreate = Event{
	Name: "InteractionCreate",
	Execute: func(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
		if interaction.Type == discordgo.InteractionApplicationCommand {
			commands.HandleSlashCommandExecute(session, interaction)
		}
		if interaction.Type == discordgo.InteractionMessageComponent {
			utils.HandleBtnExecute(session, interaction)
		}
	},
}
