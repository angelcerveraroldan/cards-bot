package api

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
)

/*
OptionsToMap

Convert the options passed in by the user into a map
"parameter": "parameter value", eg. "name": "charizard"
*/
func OptionsToMap(opts []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
	for _, option := range opts {
		optionsMap[option.Name] = option
	}

	return optionsMap
}

/*
URLtoStruct
TODO: Return an error if the card was not found
*/
func URLtoStruct(URL string, rsp any) error {
	// Get URL response
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error reading url: ", err)
		return err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error getting data from url: ", err)
		return err
	}

	// Unmarshal
	err = json.Unmarshal(responseData, &rsp)
	if err != nil {
		fmt.Println("Could not unmarshal card data: ", err)
		return err
	}

	return nil
}
