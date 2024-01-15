package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigBot struct {
	Token string    `json:"token"`
	Guild GuildInfo `json:"guild"`
}

type GuildInfo struct {
	ID      string        `json:"id"`
	Channel GuildChannels `json:"channels"`
	Roles   GuildRoles    `json:"roles"`
}

type GuildChannels struct {
	Welcome      string `json:"welcome"`
	Rules        string `json:"rules"`
	Presentation string `json:"presentation"`
}

type GuildRoles struct {
	Member    string `json:"member"`
	Validated string `json:"validated"`
}

func NewConfigBot() *ConfigBot {

	file, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier config", err)
		os.Exit(1)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	data := ConfigBot{
		Token: "",
		Guild: GuildInfo{
			ID: "",
			Channel: GuildChannels{
				Welcome:      "",
				Rules:        "",
				Presentation: "",
			},
			Roles: GuildRoles{
				Member:    "",
				Validated: "",
			},
		},
	}

	err = decoder.Decode(&data)

	if err != nil {
		fmt.Println("Erreur lors du parsing json")
		os.Exit(1)
	}

	return &data
}

var (
	ConfigBotInstance = NewConfigBot()
)
