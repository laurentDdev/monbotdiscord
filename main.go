package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"monbotdiscord/bot"
	"monbotdiscord/bot/commands"
	"monbotdiscord/bot/events"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	myBot := bot.NewBot()
	myBot.AddEvent(events.MessageCreate)
	myBot.AddEvent(events.InteractionCreate)
	myBot.AddEvent(events.MemberAdd)
	myBot.AddEvent(events.Ready)
	myBot.RegisterEvents()
	myBot.Session.Identify.Intents = discordgo.IntentsAll
	err := myBot.Session.Open()
	if err != nil {
		fmt.Println("Probleme lors du lancement du bot")
		os.Exit(1)
	}

	myBot.ApplyActivity()

	// Part Command
	myBot.AddCommand(commands.CmdPing)

	myBot.RegisterCommands()

	// Part Event
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	myBot.Session.Close()

}
