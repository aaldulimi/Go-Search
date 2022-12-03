package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func main() {
	// build the dataset
	// ScrapeNYTimes(2022, "nytimes.json")

	// build index, will be stored in memory, will later store in rocksdb
	index := buildIndex("nytimes.json")

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
