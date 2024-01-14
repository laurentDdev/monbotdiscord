package events

import (
	"fmt"
	"monbotdiscord/bot/config"
	"monbotdiscord/bot/utils"

	"github.com/bwmarrin/discordgo"
)

var Ready = Event{
	Name: "Ready",
	Execute: func(session *discordgo.Session, ready *discordgo.Ready) {
		rulesChannel, err := session.Channel(config.ConfigBotInstance.Guild.Channel.Rules)
		if err != nil {
			panic("Impossible de récupérer le channel des règles")
		}

		utils.ClearChannel(session, rulesChannel)

		embedRules := &discordgo.MessageEmbed{
			Title:       "Règles du serveur",
			Description: "Bienvenue sur notre serveur Discord dédié au développement ! Pour assurer une communauté saine et productive, veuillez respecter les règles suivantes :",
			Color:       0x00ff00,
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "1. Respect",
					Value:  "Soyez respectueux envers tous les membres du serveur. Aucun comportement abusif, harcèlement, discrimination ou trolling ne sera toléré.",
					Inline: false,
				},
				{
					Name:   "2. Contenu approprié",
					Value:  "Veuillez garder tout le contenu approprié et sûr. Aucun contenu explicite, offensant ou NSFW n'est autorisé.",
					Inline: false,
				},
				{
					Name:   "3. Aide et soutien",
					Value:  "Ce serveur est dédié à l'apprentissage et à l'entraide. N'hésitez pas à poser des questions et à aider les autres, mais ne demandez pas ou ne fournissez pas d'aide pour des activités illégales ou contraires à l'éthique.",
					Inline: false,
				},
				{
					Name:   "4. Spam et publicité",
					Value:  "Évitez le spam et la publicité non sollicitée. Ne publiez pas de liens vers d'autres serveurs Discord sans autorisation.",
					Inline: false,
				},
			},
			Footer: &discordgo.MessageEmbedFooter{
				Text: "Merci de respecter ces règles pour garder notre communauté accueillante et productive !",
			},
		}

		_, err = session.ChannelMessageSendComplex(rulesChannel.ID, &discordgo.MessageSend{
			Content: "Bienvenue sur le serveur de CodeEnsemble !",
			Embed:   embedRules,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "✅",
							},
							Label:    "J'accepte les règles",
							CustomID: string(utils.BtnRules),
						},
					},
				},
			},
		})
		if err != nil {
			fmt.Println("Erreur lors de l'envoie du message des regle", err)
		}
	},
}
