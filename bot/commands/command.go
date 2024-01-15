package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func HandleSlashCommandExecute(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	commandName := interaction.ApplicationCommandData().Name

	switch commandName {
	case "ping":
		response := CmdPingResponse(session, interaction)
		err := session.InteractionRespond(interaction.Interaction, &response)
		if err != nil {
			log.Println("Erreur lors de la réponse à la commande ping", err)
		}
	case "clear":
		response := CmdClearResponse(session, interaction)
		err := session.InteractionRespond(interaction.Interaction, &response)
		if err != nil {
			log.Println("Erreur lors de la réponse à la commande clear", err)
		}
	}
}
