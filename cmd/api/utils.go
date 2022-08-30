package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Contains
//
// Check if slice of type A contains item
func Contains[A comparable](slice []A, item A) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}

	return false
}

// ParamsToMap
//
// This function will turn a slice of strings into a map from string to string.
//
// It takes two parameters, one being the slice of strings from which we want to generate the map
// the other one being the keys for the map.
func ParamsToMap(params []string, keywords []string) map[string]string {
	paramsMap := make(map[string]string)
	currentKey := ""

	for _, s := range params {
		if Contains(keywords, s) {
			currentKey = s
			continue
		}

		if currentKey == "" {
			continue
		}

		switch paramsMap[currentKey] {
		case "":
			paramsMap[currentKey] = s
		default:
			paramsMap[currentKey] = fmt.Sprintf("%s %s", paramsMap[currentKey], s)
		}
	}

	return paramsMap
}

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

func URLEncode(s string) string {
	spaces := strings.ReplaceAll(s, " ", "%20")
	quotes := strings.ReplaceAll(spaces, "\"", "%22")
	colon := strings.ReplaceAll(quotes, ":", "%3A")

	return colon
}
