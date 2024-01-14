package events

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var MessageCreate = Event{
	Name: "MessageCreate",
	Execute: func(session *discordgo.Session, message *discordgo.MessageCreate) {

		// Ignore all messages created by the bot itself
		// This isn't required in this specific example but it's a good practice.
		if message.Author.ID == session.State.User.ID {
			return
		}
		// If the message is "ping" reply with "Pong!"
		if message.Content == "ping" {
			_, err := session.ChannelMessageSendReply(message.ChannelID, "Pong!", message.Reference())
			if err != nil {
				fmt.Println("Erreur lors de l'envoie du message", err)
				return
			}
		}

	},
}
