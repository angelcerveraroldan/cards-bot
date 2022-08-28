package api

import "fmt"

func Contains[A comparable](slice []A, item A) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}

	return false
}

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
