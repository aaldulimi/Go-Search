package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var articles []Article
	dataFile, err := os.Open("data.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(dataFile)
	json.Unmarshal(byteValue, &articles)

	for _, value := range articles {
		fmt.Println(value.Title)
	}

}
