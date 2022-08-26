package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

// id ex xy1-1
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

// Get card by name
