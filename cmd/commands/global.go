package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func helpCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Title:   "Help",
			Content: fmt.Sprintf("Help message to be implemented"),
		},
	})
}

// Heartbeat command is used to check if the bot is up and running
func heartbeat(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "I'm alive!",
		},
	})
}
