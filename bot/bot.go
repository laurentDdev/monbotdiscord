package bot

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"monbotdiscord/bot/config"
	"monbotdiscord/bot/events"
)

type Bot struct {
	Session  *discordgo.Session
	Events   []events.Event
	Commands []discordgo.ApplicationCommand
}

func NewBot() *Bot {

	discord, err := discordgo.New("Bot " + config.ConfigBotInstance.Token)
	if err != nil {
		panic("Erreur lors de la création du bot")
	}

	return &Bot{
		Session:  discord,
		Events:   make([]events.Event, 0),
		Commands: make([]discordgo.ApplicationCommand, 0),
	}
}

func (b *Bot) ApplyActivity() {
	err := b.Session.UpdateGameStatus(0, "CodeEnsemble")
	if err != nil {
		fmt.Println("Erreur lors de la mise a jour de l'activiter ", err)
	}
}

func (b *Bot) AddCommand(cmd discordgo.ApplicationCommand) {
	b.Commands = append(b.Commands, cmd)
}

func (b *Bot) RegisterCommands() {
	fmt.Println("Génération des commandes")
	fmt.Println("------------------------")

	for _, cmd := range b.Commands {
		fmt.Println("Création de la commande ", cmd.Name)
		_, err := b.Session.ApplicationCommandCreate(b.Session.State.User.ID, config.ConfigBotInstance.Guild.ID, &cmd)
		if err != nil {
			fmt.Println("Erreur lors de la création de la commande ", err)
		}
	}
	fmt.Println("------------------------")
}

func (b *Bot) AddEvent(event events.Event) {
	b.Events = append(b.Events, event)
}

func (b *Bot) RegisterEvents() {
	fmt.Println("Génération des events")
	fmt.Println("------------------------")
	for _, event := range b.Events {
		fmt.Println("Enregistrement de l'event ", event.Name)
		b.Session.AddHandler(event.Execute)
	}
	fmt.Println("------------------------")
}
