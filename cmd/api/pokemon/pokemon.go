package pokemon

/*
This file is the pokemon handler
*/

import (
	"github.com/bwmarrin/discordgo"
)

const (
	baseURL = "https://api.pokemontcg.io/v2"
)

// Commands for pokemon
var Commands = []*discordgo.ApplicationCommand{
	{
		Name:        "card-id",
		Description: "Get a card with a given id",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "id",
				Description: "id of desired card",
				Required:    true,
			},
		},
	},
	{
		Name:        "card-where",
		Description: "Get a card with given parameters, such as name: Charizard, hp: 200, etc...",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Specify cards name",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "artist",
				Description: "Specify the name of the artist that drew the cards image",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "rarity",
				Description: "Specify the rarity of the card",
				Required:    false,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "hp",
				Description: "Specify the health points of the card",
				Required:    false,
			},
		},
	},
}

// CommandsHandler for pokemon
var CommandsHandler = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"card-id":    cardId,
	"card-where": cardParams,
}
