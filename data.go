package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

var docID int = 0

type Article struct {
	Title   string `json:"title"`
	Id      int    `json:"id"`
	Summary string `json:"summary"`
	Body    string `json:"body"`
	URL     string `json:"url"`
}

func ScrapeNYTimes(year int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Write([]byte("["))
	getMonthsFromYear(file, year)
	file.Write([]byte("]"))
}

func getMonthsFromYear(file *os.File, year int) {
	yearStr := fmt.Sprintf("%d", year)
	url := "https://www.nytimes.com/sitemap/" + yearStr + "/"

	c := colly.NewCollector()
	monthLimit, monthCount := 1, 0
	c.OnHTML("ol > li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		fullLink := url + link

		if monthCount < monthLimit {
			fmt.Println(fullLink)
			getDaysFromMonth(file, fullLink)

		}
		monthCount++
	})

	c.Visit(url)
}

func getDaysFromMonth(file *os.File, url string) {
	c := colly.NewCollector()
	c.OnHTML("ol > li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		fullLink := url + link

		getArticlesFromDay(file, fullLink)
	})

	c.Visit(url)
}

func getArticlesFromDay(file *os.File, url string) {
	c := colly.NewCollector()
	c.OnHTML("#site-content > div > ul:nth-child(4) > li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		article := getArticleContent(link)

		// add article to file
		articleJSON, err := json.Marshal(article)
		if err != nil {
			panic(err)
		}

		var indentedJSON bytes.Buffer
		json.Indent(&indentedJSON, articleJSON, "", "\t")

		file.Write(indentedJSON.Bytes())
		file.Write([]byte(","))

	})

	c.Visit(url)

}

func getArticleContent(url string) Article {
	article := Article{URL: url}
	c := colly.NewCollector()
	c.OnHTML("#story > header > div:nth-child(3) > h1", func(e *colly.HTMLElement) {
		title := e.Text
		article.Title = title
	})

	c.OnHTML("#article-summary", func(e *colly.HTMLElement) {
		summary := e.Text
		article.Summary = summary
	})

	c.OnHTML("#story > section > div > div > p", func(e *colly.HTMLElement) {
		article.Body += e.Text

	})
	c.Visit(url)

	article.Id = docID
	docID++

	return article

}
