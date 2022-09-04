package pokemon

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

/*
Here are the methods that turn a card into a response (like a message, an embed, etc ...)
*/

func (c Card) sendInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Finding card...",
			Embeds: []*discordgo.MessageEmbed{
				c.toEmbed(),
			},
			Title: "Card interaction",
		},
	})
}

/*
Turn a card into an embed
*/
func (c Card) cardData() string {
	return fmt.Sprintf(">>> **Name:** %s\n**Id:** %s\n**Hp:** %s\n**Artist:** %s\n**Type/s:** %s",
		c.Name,
		c.Id,
		c.Hp,
		c.Artist,
		strings.Join(c.Types, ", "))
}

func (c Card) setData() string {
	return fmt.Sprintf(">>> **Name:** %s\n**Id:** %s", c.Set.Name, c.Set.Id)
}

// sendEmbed
//
// Method for card that turns the cards data into an embed, and send it to the channel id in m.ChannelID
func (c Card) toEmbed() *discordgo.MessageEmbed {
	fields := []*discordgo.MessageEmbedField{
		{
			Name:   "Card Data:",
			Value:  c.cardData(),
			Inline: false,
		},
		{
			Name:   "Set Data:",
			Value:  c.setData(),
			Inline: false,
		},
	}

	return &discordgo.MessageEmbed{
		Title:       "Pokemon card",
		Description: "Information on the requested card",
		Color:       0x0099FF,
		Image:       &discordgo.MessageEmbedImage{URL: c.Images.Large},
		Fields:      fields,
	}
}
