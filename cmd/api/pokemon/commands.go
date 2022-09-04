package pokemon

import (
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/api"
	"github.com/bwmarrin/discordgo"
)

/*
Here go the pokemon commands
All the functions here MUST have the signature func(s *discordgo.Session, i *discordgo.InteractionCreate)
*/

/*
cardId this function is used to get a card by id

Every card has a different and unique id, so pagination isn't required
The only Option is the "id" field, this is a required, string.
*/
func cardId(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// The required option id should be used to find the card data
	optionsMap := api.OptionsToMap(i.ApplicationCommandData().Options)

	id := optionsMap["id"].StringValue()

	card, err := getCardById(id)
	if err != nil {
		fmt.Printf("There was an error calling getCardById(\"%s\")\n", id)
		return
	}

	card.sendInteraction(s, i)
}

func cardParams(s *discordgo.Session, i *discordgo.InteractionCreate) {
	optionsMap := api.OptionsToMap(i.ApplicationCommandData().Options)

	// TODO: Remove this at some point
	fmt.Println("Card by params requested with:")
	for k, v := range optionsMap {
		fmt.Printf("    %v: %v", k, v)
	}
	fmt.Println("")

	// Get card from the parameters passed in as options

	// If the user passed no parameters in, send an error message
	if len(optionsMap) == 0 {
		// TODO: Send error message
		return
	}

	// Turn map into map[string]string
	// getCardByParams expect map[string]string as a parameter
	parameters := make(map[string]string)

	for k, v := range optionsMap {
		parameters[k] = v.StringValue()
	}

	cards, _ := getCardsByParams(parameters)
	if len(cards) == 0 {
		// TODO: Send error message
		return
	}
	cards[0].sendInteraction(s, i)
}
