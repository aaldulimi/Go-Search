package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func Search(searchQuery string, searchLimit int, index map[string][]int, filename string) []Article {
	searchTokens := Tokenize(searchQuery)

	// hashmap to store all doc ids where token exists, for each token
	searchDocs := buildDocsMap(searchTokens, index)
	// list of Articles that matched the search query
	results := docSearcher(filename, searchDocs, len(searchTokens), searchLimit)

	return results
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
