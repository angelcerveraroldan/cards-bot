package main

import (
	"flag"
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/commands"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

const prefix = "!h"

var discordToken string

func init() {
	flag.StringVar(&discordToken, "t", "", "Discord discordToken")

	flag.Parse()

	if discordToken == "" {
		panic("Discord token needed")
	}
}

func main() {
	ds, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("Error creating discord session, err: ", err)
	}

	// This function will be run every time a message is sent into any channel that the commands can read
	ds.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		messages := strings.Fields(m.Message.Content)

		if messages[0] == prefix {
			commands.RunCommand(messages[1:], s, m)
		}
	})

	err = ds.Open()
	if err != nil {
		fmt.Println("Error opening discord commands: ", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	ds.Close()
}
