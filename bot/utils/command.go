package utils

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"monbotdiscord/bot/commands"
)

func HandleSlashCommandExecute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	commandName := interaction.ApplicationCommandData().Name

	switch commandName {
	case "ping":
		response := commands.CmdPingResponse(interaction)
		err := session.InteractionRespond(interaction.Interaction, &response)
		if err != nil {
			log.Println("Erreur lors de la réponse à la commande ping", err)
		}
	}
}
