package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func Scrape() {
	// Initialize the collector
	c := colly.NewCollector(
		colly.AllowedDomains("hackerspaces.org", "wiki.hackerspaces.org"),
	)

	// Extract links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))

		// Filter out irrelevant links (modify if needed)
		if strings.Contains(link, "hackerspaces.org/wiki/") {
			fmt.Println("Visiting:", link)
			c.Visit(link) // Visit each extracted link
		}
	})

	// Extract article content from visited pages
	c.OnHTML("p", func(e *colly.HTMLElement) {
		// Print or store the text content of <p> elements (blog/article content)
		fmt.Println(e.Text)
	})

	// Handle errors
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Error:", err)
	})

	// Start scraping from the main page
	c.Visit("https://hackerspaces.org/")
}
