package pokemon

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	baseURL = "https://api.pokemontcg.io/v2"
)

var (
	// TODO: Add params such as attack.name or set.id
	searchKeys = []string{"name", "subtype", "hp"}
)

// GetCardData -> Use appropriate function to get card data
func GetCardData(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	getMethod := args[0]

	switch getMethod {
	case "id":
		card, err := getCardById(args[1])
		if err != nil {
			fmt.Println("Could not find card")
			break
		}

		card.printCardData(s, m)
	}

}

func (card Card) otherInformation() string {
	var level string
	level = card.Level
	if level == "" {
		level = "N/A"
	}
	return fmt.Sprintf(" - Supertype: %s\n - Level: %s\n - Hp: %s\n - Type/s: %s", card.Supertype, level, card.Hp, strings.Join(card.Types, ", "))
}

// URGENT TODO: Make this function print a prettier message, make this a method on type card?
func (card Card) printCardData(s *discordgo.Session, m *discordgo.MessageCreate) {
	if card.Name == "" {
		printError("Couldn't find card", s, m)
		return
	}

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Name",
			Value:  card.Name,
			Inline: false,
		},
		{
			Name:   "Card ID",
			Value:  card.Id,
			Inline: true,
		},
		{
			Name:   "Set ID",
			Value:  card.Set.Id,
			Inline: true,
		},
		{
			Name:   "Set Name",
			Value:  card.Set.Name,
			Inline: true,
		},
		{
			Name:   "Other information",
			Value:  card.otherInformation(),
			Inline: false,
		},
	}

	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Pokemon card",
		Description: "Information on the requested card",
		Color:       0x0099FF,
		Image: &discordgo.MessageEmbedImage{
			URL: card.Images.Small,
		},
		Fields: fields,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
}

func printError(err string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("There was an error: %s", err))
}
