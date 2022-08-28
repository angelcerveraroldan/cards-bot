package pokemon

import (
	"encoding/json"
	"fmt"
	"github.com/angelcerveraroldan/cards-bot/cmd/api"
	"io"
	"net/http"
	"strings"
)

func urlResponse(URL string) ([]byte, error) {
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("Error reading url: ", err)
		return []byte{}, err
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error getting data from url: ", err)
		return []byte{}, err
	}

	return responseData, nil
}

// C is either a response with a card, or an array of cards
func unmarshal[C Response](data []byte, rsp *C) error {
	err := json.Unmarshal(data, rsp)
	if err != nil {
		fmt.Println("Could not unmarshal card data: ", err)
		return err
	}

	return nil
}

// getCardById
//
// Each card has a unique id, given that id, find the card
func getCardById(id string) (Card, error) {
	URL := baseURL + "/cards/" + id

	responseData, err := urlResponse(URL)
	if err != nil {
		return Card{}, err
	}

	var rsp CardResponse
	err = unmarshal(responseData, &rsp)
	if err != nil {
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
		paramsStr = append(paramsStr, fmt.Sprintf("%s:\"%s\"", k, v))
	}

	URL = fmt.Sprintf("%s%s", URL, URLEncode(strings.Join(paramsStr, " ")))

	responseData, err := urlResponse(URL)
	if err != nil {
		return []Card{}, err
	}

	var rsp CardsResponse
	err = unmarshal(responseData, &rsp)
	if err != nil {
		return []Card{}, err
	}

	return rsp.Cards, nil
}

func URLEncode(s string) string {
	spaces := strings.ReplaceAll(s, " ", "%20")
	quotes := strings.ReplaceAll(spaces, "\"", "%22")
	colon := strings.ReplaceAll(quotes, ":", "%3A")

	return colon
}
