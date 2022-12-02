package main

import (
	"strings"
)

func SplitText(text string) []string {
	splitText := strings.Split(text, " ")
	return splitText
}

func ConvertToLowercase(tokens []string) []string {
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}

	return tokens
}
