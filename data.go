package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	nytimes := "https://www.nytimes.com/sitemap/2021/"
	c := colly.NewCollector()

	c.OnHTML("ol > li", func(e *colly.HTMLElement) {
		link := e.ChildAttr("a", "href")
		fullLink := nytimes + link
		fmt.Println(fullLink)
	})

	c.Visit(nytimes)

}
