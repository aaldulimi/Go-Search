package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	searchQuery := "Blackberry phones end of life"
	searchLimit := 5
	searchTokens := Tokenize(searchQuery)

	filename := "nytimes.json"
	// build the dataset
	// ScrapeNYTimes(2022, filename)

	// build index, will be stored in memory, will later store in rocksdb
	index := buildIndex(filename)

	// hashmap to store all doc ids where token exists, for each token
	searchDocs := buildDocsMap(searchTokens, index)
	// list of Articles that matched the search query
	results := docSearcher(filename, searchDocs, len(searchTokens), searchLimit)

	printResults(results, true, false, false, false)

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

func buildDocsMap(searchTokens []string, index map[string][]int) map[int]int {
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

	return searchDocs
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

func printResults(results []Article, title bool, summary bool, body bool, url bool) {
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
