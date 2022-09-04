package main

import (
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/commands"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var discordToken string
var ds *discordgo.Session

func init() {
	// Load discord token from docker env variables
	discordToken = os.Getenv("TOKEN")

	if discordToken == "" {
		panic("Discord token needed")
	}

	var err error
	ds, err = discordgo.New("Bot " + discordToken)
	if err != nil {
		fmt.Println("Error creating discord session, err: ", err)
	}
}

func main() {
	err := ds.Open()
	if err != nil {
		fmt.Println("Error opening discord commands: ", err)
		return
	}

	log.Println("Adding commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands.AllCommands))
	for i, v := range commands.AllCommands {
		// guildID "" will make the commands global
		cmd, err := ds.ApplicationCommandCreate(ds.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}

		registeredCommands[i] = cmd
	}

	// Handle interactions (This will be executed every time a slash-command is used)
	ds.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// TODO: rename h
		if h, ok := commands.AllHandlers[i.ApplicationCommandData().Name]; ok {
			// Keep a log of all commands that are executed
			fmt.Printf("Command %s is being run.\n", i.ApplicationCommandData().Name)
			h(s, i)
		}
	})

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	ds.Close()
}
