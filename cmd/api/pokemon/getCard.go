package pokemon

import (
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/api"
	"strings"
)

// getCardById
//
// Each card has a unique id, given that id, find the card
func getCardById(id string) (Card, error) {
	URL := baseURL + "/cards/" + id

	var rsp CardResponse
	err := api.URLtoStruct(URL, &rsp)
	if err != nil {
		fmt.Println("Error unmarshaling card data")
		return Card{}, err
	}

	return rsp.Card, nil
}

// getCardByParams
//
// Get an array of cards that match a set of parameters
// Example parameter: name charizard
// ^^ This will return an array containing all charizard cards
func getCardsByParams(params []string) ([]Card, error) {
	URL := baseURL + "/cards?q="

	paramsMap := api.ParamsToMap(params, searchKeys)

	var paramsStr []string
	for k, v := range paramsMap {
		// ignore empty keys
		if v != "" {
			paramsStr = append(paramsStr, fmt.Sprintf("%s:\"%s\"", k, v))
		}
	}

	URL = fmt.Sprintf("%s%s", URL, URLEncode(strings.Join(paramsStr, " ")))

	var rsp CardsResponse
	err := api.URLtoStruct(URL, &rsp)
	if err != nil {
		fmt.Println("Error unmarshaling card data")
		return nil, err
	}

	return rsp.Cards, nil
}

func URLEncode(s string) string {
	spaces := strings.ReplaceAll(s, " ", "%20")
	quotes := strings.ReplaceAll(spaces, "\"", "%22")
	colon := strings.ReplaceAll(quotes, ":", "%3A")

	return colon
}
