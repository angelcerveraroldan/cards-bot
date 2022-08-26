package pokemon

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

const (
	baseURL = "https://api.pokemontcg.io/v2"
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

func (card Card) ToMessage() string {
	return fmt.Sprintf("Name: %s\nId: %s\nHp: %s\nTypes: %v", card.Name, card.Id, card.Hp, card.Types)
}

// URGENT TODO: Make this function print a prettier message, make this a method on type card?
func (card Card) printCardData(s *discordgo.Session, m *discordgo.MessageCreate) {
	if card.Name == "" {
		printError("Couldn't find card", s, m)
		return
	}

	s.ChannelMessageSend(m.ChannelID, card.ToMessage())
}

func printError(err string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("There was an error: %s", err))
}
