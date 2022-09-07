package pokemon

/*
This file will be used to gather data from the pokemon api

Every function here, should take search parameters as params, and return []Card or Card (along with error)
*/

import (
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/api"
	"strings"
)

/*
getCardById

Each card has a unique id, given that id, find the card
*/
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

/*
getCardByParams

get a card that matches certain parameters
*/
func getCardsByParams(params map[string]string) ([]Card, error) {
	URL := baseURL + "/cards?q="

	var paramsStr []string
	for k, v := range params {
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

/*
URLEncode

When making the query, certain characters need to be encoded, for example ' ' should be %20 or +
*/
func URLEncode(s string) string {
	spaces := strings.ReplaceAll(s, " ", "%20")
	quotes := strings.ReplaceAll(spaces, "\"", "%22")

	return quotes
}
