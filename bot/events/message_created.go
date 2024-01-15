package events

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"monbotdiscord/bot/config"
	"time"
)

var MessageCreate = Event{
	Name: "MessageCreate",
	Execute: func(session *discordgo.Session, message *discordgo.MessageCreate) {

		if message.Author.ID == session.State.User.ID {
			return
		}

		if message.ChannelID == config.ConfigBotInstance.Guild.Channel.Presentation {
			err := session.ChannelMessageDelete(message.ChannelID, message.ID)

			if err != nil {
				log.Println("Erreur lors de la suppression du message", err)
			}

			embed := &discordgo.MessageEmbed{
				Thumbnail: &discordgo.MessageEmbedThumbnail{
					URL: message.Author.AvatarURL(""),
				},
				Title:       "Merci de t'être présenté(e) !",
				Description: message.Content,
				Color:       0x00ff00,
				Timestamp:   time.Now().Format(time.RFC3339),
			}
			msgSend, err := session.ChannelMessageSendEmbed(message.ChannelID, embed)
			if err != nil {
				log.Println("Erreur lors de l'envoi du message", err)
			}

			err = session.MessageReactionAdd(message.ChannelID, msgSend.ID, "✅")
			if err != nil {
				log.Println("Erreur lors de l'ajout de la réaction", err)
			}

		}

	},
}
