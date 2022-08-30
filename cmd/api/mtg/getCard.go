package mtg

import (
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/api"
	"strings"
)

// getCardByParams
//
// Get an array of cards that match a set of parameters
// Example parameter: name charizard
// ^^ This will return an array containing all charizard cards
func getCardsByParams(params []string) ([]Card, error) {
	URL := baseURL + "/cards?"

	paramsMap := api.ParamsToMap(params, searchKeys)

	var paramsStr []string
	for k, v := range paramsMap {

		if strings.Contains(v, " AND ") {
			v = strings.ReplaceAll(v, " AND ", ",")
		}

		// ignore empty keys
		if v != "" {
			paramsStr = append(paramsStr, fmt.Sprintf("%s=\"%s\"", k, v))
		}
	}

	URL = fmt.Sprintf("%s%s", URL, api.URLEncode(strings.Join(paramsStr, "&")))
	fmt.Println(URL)

	var rsp CardsResponse
	err := api.URLtoStruct(URL, &rsp)
	if err != nil {
		fmt.Println("Error unmarshaling card data")
		return nil, err
	}

	return rsp.Cards, nil
}
