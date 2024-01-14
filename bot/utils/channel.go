package utils

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func ClearChannel(session *discordgo.Session, channel *discordgo.Channel) {
	messages, err := session.ChannelMessages(channel.ID, 100, "", "", "")
	if err != nil {
		log.Println("Erreur lors de la récupération des messages du channel", err)
		return
	}

	for _, message := range messages {
		err = session.ChannelMessageDelete(channel.ID, message.ID)
		if err != nil {
			log.Println("Erreur lors de la suppression du message", err)
			return
		}
	}
}
