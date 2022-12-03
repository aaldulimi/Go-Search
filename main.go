package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// build the dataset
	// ScrapeNYTimes(2022, "nytimes.json")

	// build index, will be stored in memory, will later store in rocksdb
	index := buildIndex("nytimes.json")

	searchQuery := "Donald"
	searchTokens := Tokenize(searchQuery)

	// use hashmap to store list of ids rather than slice, prevents iterating over array
	searchDocs := make(map[int]int)

	if len(searchTokens) == 1 {
		token := searchTokens[0]
		docsSlice := index[token]

		for _, docId := range docsSlice {
			searchDocs[docId] = 1
		}

		results := docSearcher("nytimes.json", searchDocs)
		for _, article := range results {
			fmt.Println(article.Title)
		}
	}

}

func buildIndex(filename string) map[string][]int {
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

func docSearcher(filename string, docsId map[int]int) []Article {
	var articles []Article
	var returnArticles []Article

	dataFile, _ := os.Open(filename)
	byteValue, _ := ioutil.ReadAll(dataFile)
	json.Unmarshal(byteValue, &articles)

	for _, article := range articles {
		if _, ok := docsId[article.Id]; ok {
			returnArticles = append(returnArticles, article)
		}
	}

	return returnArticles
}
