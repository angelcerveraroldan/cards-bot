package pokemon

import (
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/messages"
	"github.com/bwmarrin/discordgo"
	"strings"
)

const (
	baseURL = "https://api.pokemontcg.io/v2"
)

var (
	searchKeys = []string{"name", "subtypes", "hp", "types", "set.name", "set.id", "attacks.name", "artist", "rarity"}
)

// GetCardData
//
// Pokemon API call handler
// Use appropriate function to get card data
func GetCardData(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	if len(args) < 1 {
		// TODO: Implement helper message here
		return
	}

	getMethod := args[0]

	switch getMethod {
	case "id":
		card, err := getCardById(args[1])
		if err != nil {
			fmt.Println("Could not find card, err: ", err)
			messages.Error(s, m, fmt.Sprintf("Could not find card with id: %s", args[1]))
			return
		}

		card.sendEmbed(s, m)

	case "where":
		cards, err := getCardsByParams(args[1:])
		if err != nil || len(cards) == 0 {
			fmt.Println("Could not find card, err: ", err)
			messages.Error(s, m, "Could not find card with given parameters")
			return
		}

		if len(cards) > 1 {
			messages.Send(s, m, fmt.Sprintf("There multiple cards with the requested parameters, showing one of them"))
		}

		cards[0].sendEmbed(s, m)
	}
}

func (card Card) cardData() string {
	return fmt.Sprintf(">>> **Name:** %s\n**Id:** %s\n**Hp:** %s\n**Artist:** %s\n**Type/s:** %s",
		card.Name,
		card.Id,
		card.Hp,
		card.Artist,
		strings.Join(card.Types, ", "))

}

func (card Card) setData() string {
	return fmt.Sprintf(">>> **Name:** %s\n**Id:** %s", card.Set.Name, card.Set.Id)
}

// sendEmbed
//
// Method for card that turns the cards data into an embed, and send it to the channel id in m.ChannelID
func (card Card) sendEmbed(s *discordgo.Session, m *discordgo.MessageCreate) {
	if card.Name == "" {
		messages.Error(s, m, "Couldn't find card")
		return
	}

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Card Data:",
			Value:  card.cardData(),
			Inline: false,
		},
		{
			Name:   "Set Data:",
			Value:  card.setData(),
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
