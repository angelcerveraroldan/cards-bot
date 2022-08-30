package messages

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Send(s *discordgo.Session, m *discordgo.MessageCreate, message string) {
	_, err := s.ChannelMessageSend(m.ChannelID, message)
	if err != nil {
		fmt.Println("There has been an error sending message")
		return
	}
}

func Error(s *discordgo.Session, m *discordgo.MessageCreate, error string) {
	Send(s, m, fmt.Sprintf("There has been an error: %s", error))
}

func SendEmbed(s *discordgo.Session, m *discordgo.MessageCreate, embed *discordgo.MessageEmbed) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, embed)
	if err != nil {
		Error(s, m, "Couldn't send embed")
		fmt.Println("There has been an error sending embed")
		return
	}
}
