package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := "nytimes.json"
	// build the dataset
	// ScrapeNYTimes(2022, filename)

	// build index, will be stored in memory, will later store in rocksdb
	index := buildIndex(filename)

	searchQuery := "Blackberry"
	searchLimit := 10
	searchTokens := Tokenize(searchQuery)

	// use hashmap to store list of ids rather than slice, prevents iterating over array
	searchDocs := make(map[int]int)

	for _, token := range searchTokens {
		docsSlice := index[token]

		for _, docId := range docsSlice {
			if _, ok := searchDocs[docID]; ok {
				searchDocs[docId]++
			} else {
				searchDocs[docId] = 1
			}
		}
	}

	results := docSearcher(filename, searchDocs, len(searchTokens), searchLimit)
	for _, article := range results {
		fmt.Println(article.Title)
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

func docSearcher(filename string, searchDocs map[int]int, rank int, limit int) []Article {
	var articles []Article
	var returnArticles []Article
	currentRank := rank

	dataFile, _ := os.Open(filename)
	byteValue, _ := ioutil.ReadAll(dataFile)
	json.Unmarshal(byteValue, &articles)

	for len(returnArticles) < limit {
		docIterator(articles, searchDocs, &returnArticles, currentRank, limit)
		currentRank--

		if currentRank == 0 {
			break
		}
	}

	return returnArticles

}

func docIterator(articles []Article, searchDocs map[int]int, returnArticles *[]Article, currentRank int, limit int) {
	articleCount := 0

	for _, article := range articles {
		if val, ok := searchDocs[article.Id]; ok && val == currentRank {
			if article.Title != "" {
				*returnArticles = append(*returnArticles, article)
				articleCount++
			}
		}

		if articleCount == limit {
			break
		}
	}
}
