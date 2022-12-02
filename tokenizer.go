package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/kljensen/snowball"
)

func main() {
	testString := "this is A STRING, THAT CONTAINS. SOME Words, running."

	fmt.Println("Split text:")
	splitText := SplitText(testString)
	for _, value := range splitText {
		fmt.Print(value, " ")
	}

	fmt.Println("\nConvert to lowercase:")
	lowerCase := ConvertToLowercase(splitText)
	for _, value := range lowerCase {
		fmt.Print(value, " ")
	}

	fmt.Println("\nRemove punctuation:")
	noPunc := RemovePunctuation(lowerCase)
	for _, value := range noPunc {
		fmt.Print(value, " ")
	}

	fmt.Println("\nRemove stop words:")
	noStopWords := removeStopWords(noPunc)
	for _, value := range noStopWords {
		fmt.Print(value, " ")
	}

	fmt.Println("\nKeep only stem words:")
	onlyStem := KeepStem(noStopWords)
	for _, value := range onlyStem {
		fmt.Print(value, " ")
	}
}

// split a string to a slice of tokens
func SplitText(text string) []string {
	splitText := strings.Split(text, " ")
	return splitText
}

// convert all strings in a slice to lowercase
func ConvertToLowercase(tokens []string) []string {
	for i, token := range tokens {
		tokens[i] = strings.ToLower(token)
	}

	return tokens
}

// remove punctuation, keep only letters, numbers and underscores
func RemovePunctuation(tokens []string) []string {
	r, err := regexp.Compile("[^\\w\\s]")
	if err != nil {
		panic(err)
	}

	for i, token := range tokens {
		tokens[i] = r.ReplaceAllString(token, "")
	}

	return tokens
}

// remove stop words, i.e. most common 25 words in english language
func removeStopWords(tokens []string) []string {
	stopWords := []string{"the", "be", "to", "of", "and", "a", "in", "that", "have",
		"i", "it", "for", "not", "on", "with", "he", "as", "you",
		"do", "at", "this", "but", "his", "by", "from"}

	stopWordsHash := make(map[string]int)
	for _, value := range stopWords {
		stopWordsHash[value] = 1
	}

	var newSlice []string
	for _, value := range tokens {
		_, ok := stopWordsHash[value]

		if !ok {
			newSlice = append(newSlice, value)
		}
	}

	return newSlice
}

func KeepStem(tokens []string) []string {
	var stemWords []string
	for _, token := range tokens {
		stemmed, err := snowball.Stem(token, "english", true)

		if err == nil {
			stemWords = append(stemWords, stemmed)
		}

	}

	return stemWords
}
