package api

import (
	"github.com/bwmarrin/discordgo"
	"testing"
)

func TestOptionsToMap(t *testing.T) {
	arr := []*discordgo.ApplicationCommandInteractionDataOption{
		{
			Name:  "name",
			Type:  discordgo.ApplicationCommandOptionString,
			Value: "values",
		},
	}

	got := OptionsToMap(arr)

	if got["name"] != arr[0] {
		t.Errorf("The option did not match its name\n")
	}
}
