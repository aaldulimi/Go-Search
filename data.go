package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type article struct {
	title   string
	summary string
	body    string
	url     string
}

func main() {
	getMonthsFromYear(2021)

}

func getMonthsFromYear(year int) {
	yearStr := fmt.Sprintf("%d", year)
	url := "https://www.nytimes.com/sitemap/" + yearStr + "/"

	c := colly.NewCollector()
	c.OnHTML("ol > li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		fullLink := url + link

		getDaysFromMonth(fullLink)
	})

	c.Visit(url)
}

func getDaysFromMonth(url string) {
	c := colly.NewCollector()
	c.OnHTML("ol > li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		fullLink := url + link

		getArticlesFromDay(fullLink)
	})

	c.Visit(url)
}

func getArticlesFromDay(url string) {
	c := colly.NewCollector()
	c.OnHTML("#site-content > div > ul:nth-child(4) > li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		article := getArticleContent(link)
		// add article to file

	})

	c.Visit(url)

}

func getArticleContent(url string) article {
	article := article{url: url}
	c := colly.NewCollector()
	c.OnHTML("#story > header > div:nth-child(3) > h1", func(e *colly.HTMLElement) {
		title := e.Text
		article.title = title
	})

	c.OnHTML("#article-summary", func(e *colly.HTMLElement) {
		summary := e.Text
		article.summary = summary
	})

	c.OnHTML("#story > section > div > div > p", func(e *colly.HTMLElement) {
		article.body += e.Text

	})
	c.Visit(url)

	return article

}
