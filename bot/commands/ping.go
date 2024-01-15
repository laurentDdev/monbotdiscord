package commands

import "github.com/bwmarrin/discordgo"

var CmdPing = discordgo.ApplicationCommand{
	Name:        "ping",
	Description: "Renvoie pong",
}

func CmdPingResponse(session *discordgo.Session, interaction *discordgo.InteractionCreate) discordgo.InteractionResponse {
	response := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!",
		},
	}

	return response
}
