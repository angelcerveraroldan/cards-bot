package commands

import (
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/api/pokemon"
	"github.com/angelcerveraroldan/cards-bot/cmd/messages"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// RunCommand
//
// Check the command the user sent, and redirect to the right function
func RunCommand(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	command := strings.ToLower(args[0])

	if len(args) < 2 && command != "help" {
		confused(s, m)
		return
	}

	switch command {
	case "h", "help":
		help(s, m)
	case "heartbeat":
		heartBeat(s, m)
	case "pkm", "pokemon":
		pokemon.GetCardData(args[1:], s, m)

	default:
		confused(s, m)
	}

}

func heartBeat(s *discordgo.Session, m *discordgo.MessageCreate) {
	messages.Send(s, m, "I'm alive!")

}

func help(s *discordgo.Session, m *discordgo.MessageCreate) {
	messages.SendEmbed(s, m, &discordgo.MessageEmbed{
		Title:       "Help",
		Description: "Here are the commands you can use",
		Color:       0x0099FF,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name: "Pokemon Commands",
				Value: fmt.Sprintf(`
To get information on a pokemon card, run "!h pokemon" followed by the name of the command you want to run, and its parameters.

There are two commands:

**id**
This command is used to get a card's information from its unique id. Example: !h pokemon id xy1-2

**where**
This command is used to get a card with certain parameters. 
The available parameter are the cards    name, subtypes, hp, types, set.name, set.id, attacks.name, artist, rarity
Example: !h pokemon where name charizard rarity rare
`),
				Inline: false,
			},
			{
				Name:   "Magic the gathering Commands",
				Value:  "Coming Soon",
				Inline: false,
			},
		},
	})
}

func confused(s *discordgo.Session, m *discordgo.MessageCreate) {
	messages.Send(s, m, "This is not a valid command, run !h help to see all valid commands")
}
