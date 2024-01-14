package events

import (
	"fmt"
	"monbotdiscord/bot/config"
	"time"

	"github.com/bwmarrin/discordgo"
)

var MemberAdd = Event{
	Name: "MemberAdd",
	Execute: func(s *discordgo.Session, m *discordgo.GuildMemberAdd) {
		welcomeChannel, err := s.Channel(config.ConfigBotInstance.Guild.Channel.Welcome)
		if err != nil {
			fmt.Println("Erreur lors de la récupération du channel de bienvenue", err)
			return
		}
		embedWelcome := &discordgo.MessageEmbed{
			Title:       "Bienvenue sur " + s.State.User.Username + " !",
			Description: "Nous te souhaitons la bienvenue sur le serveur " + m.Member.User.Username + " !",
			Color:       0x00ff00,
			Thumbnail: &discordgo.MessageEmbedThumbnail{
				URL: m.User.AvatarURL(""),
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "Règles",
					Value:  "N'oubliez pas de lire les règles du serveur dans le canal #règles.",
					Inline: false,
				},
				{
					Name:   "Présentation",
					Value:  "Nous aimerions en savoir plus sur vous ! Présentez-vous dans le canal #présentations.",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Amusez-vous bien !",
			},
			Timestamp: time.Now().Format(time.RFC3339),
		}

		err = s.GuildMemberRoleAdd(m.GuildID, m.Member.User.ID, config.ConfigBotInstance.Guild.Roles.Member)
		_, err = s.ChannelMessageSendEmbed(welcomeChannel.ID, embedWelcome)
		if err != nil {
			fmt.Println("Erreur lors de l'envoie du message de bienvenue", err)
			return
		}
	},
}
