package commands

import (
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/api/pokemon"
	"github.com/bwmarrin/discordgo"
)

var defaultCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "help",
		Description: "This command will explain how to use the bot",
	},
	{
		Name:        "heartbeat",
		Description: "This command is used to check if the bot is active and working",
	},
}

func getAllCommands() []*discordgo.ApplicationCommand {
	return append(pokemon.Commands, defaultCommands...)
}

var AllCommands = getAllCommands()

// Base bot commands
var defaultHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"heartbeat": heartbeat,
	"help":      helpCommand,
}

// CommandHandlers returns all commandHandlers
func getAllCommandHandlers() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Add pokemon commands
	for k, v := range pokemon.CommandsHandler {
		if defaultHandlers[k] != nil {
			fmt.Println("Overwriting command: " + k)
		}

		defaultHandlers[k] = v
	}

	return defaultHandlers
}

var AllHandlers = getAllCommandHandlers()
