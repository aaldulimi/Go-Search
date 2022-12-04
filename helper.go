package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func BuildIndex(filename string) map[string][]int {
	var articles []Article
	index := make(map[string][]int)

	dataFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(dataFile)
	json.Unmarshal(byteValue, &articles)

	// index all the articles
	for _, article := range articles {
		tokens := Tokenize(article.Body)

		for _, token := range tokens {
			if val, ok := index[token]; ok {
				// token exists in index, add index id to slice
				index[token] = append(val, article.Id)
			} else {
				// crate slice to add token
				index[token] = []int{article.Id}
			}
		}
	}

	return index
}

func PrintResults(results []Article, title bool, summary bool, body bool, url bool) {
	for _, article := range results {
		if title {
			fmt.Println(article.Title)
		}

		if summary {
			fmt.Println(article.Summary)
		}

		if body {
			fmt.Println(article.Body)
		}

		if url {
			fmt.Println(article.URL)
		}

	}
}
