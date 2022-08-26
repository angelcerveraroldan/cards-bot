package pokemon

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
)

const (
	baseURL = "https://api.pokemontcg.io/v2"
)

func GetCardData(args []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	getMethod := args[0]

	switch getMethod {
	case "id":
		card, err := getCardById(args[1])
		if err != nil {
			fmt.Println("Could not find card")
			break
		}

		if card.Name == "" {
			printError("Could not find card", s, m)
			break
		}

		printCardData(card, s, m)
	}

}

// id ex xy1-1
func getCardById(id string) (Card, error) {
	URL := baseURL + "/cards/" + id

	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error reading url: ", err)
		return Card{}, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error getting data from url: ", err)
		return Card{}, err
	}

	var rsp Response
	err = json.Unmarshal(responseData, &rsp)
	if err != nil {
		fmt.Println("Could not unmarshal")
		return Card{}, err
	}

	return rsp.Data, nil
}

// URGENT TODO: Make this function print a prettier message, make this a method on type card?
func printCardData(card Card, s *discordgo.Session, m *discordgo.MessageCreate) {
	msg := fmt.Sprintf("Name: %s\nId: %s\nHp: %s\nTypes: %v", card.Name, card.Id, card.Hp, card.Types)

	s.ChannelMessageSend(m.ChannelID, msg)
}

func printError(err string, s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("There was an error: %s", err))
}
