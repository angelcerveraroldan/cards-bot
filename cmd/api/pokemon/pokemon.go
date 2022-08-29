package pokemon

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"math/rand"
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
	getMethod := args[0]

	switch getMethod {
	case "id":
		card, err := getCardById(args[1])
		if err != nil {
			fmt.Println("Could not find card, err: ", err)
			printError(fmt.Sprintf("Could not find card with id: %s", args[1]), s, m)
			break
		}

		card.printCardData(s, m)

	case "where":
		cards, err := getCardsByParams(args[1:])
		if err != nil || len(cards) == 0 {
			fmt.Println("Could not find card, err: ", err)
			printError("Could not find card with given parameters", s, m)
			break
		}

		if len(cards) == 250 {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("There were 250+ cards with the requested parameters, showing one of them"))
		} else if len(cards) > 1 {
			s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("There were %d cards with the requested parameters, showing one of them", len(cards)))
		}

		cards[rand.Intn(len(cards))].printCardData(s, m)
	}
}

func (card Card) dataToShow() struct {
	setData  string
	cardData string
} {
	setData := fmt.Sprintf(">>> **Name**: %s\n**Id**: %s", card.Set.Name, card.Set.Id)
	cardData := fmt.Sprintf(">>> **Name**: %s\n**Id**: %s\n**Hp**: %s\n**Artist**: %s\n**Type/s**: %s",
		card.Name,
		card.Id, card.Hp,
		card.Artist,
		strings.Join(card.Types, ", "))
	return struct {
		setData  string
		cardData string
	}{setData, cardData}
}

// printCardData
//
// Method for card that turns the cards data into an embed, and send it to the channel id in m.ChannelID
func (card Card) printCardData(s *discordgo.Session, m *discordgo.MessageCreate) {
	if card.Name == "" {
		printError("Couldn't find card", s, m)
		return
	}

	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Card Data:",
			Value:  card.dataToShow().cardData,
			Inline: false,
		},
		{
			Name:   "Set Data:",
			Value:  card.dataToShow().setData,
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

// printError
//
// When there is an error, send a message into the chat informing the user/s of said error
func printError(err string, s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err2 := s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("There was an error: %s", err))
	if err2 != nil {
		fmt.Println("Error sending message")
		return
	}
}
