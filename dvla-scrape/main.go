package main

import "fmt"
import "github.com/gocolly/colly/v2"

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnHTML(`table`, func(e *colly.HTMLElement) {
		e.ForEach(`tbody`, func(_ int, body *colly.HTMLElement) {
			body.ForEach(`tr`, func(_ int, row *colly.HTMLElement) {
				head := row.ChildText(`th`)
				if head == "Apply for a first provisional driving licence with UK identity" {
					fmt.Println("Currently Processing Date:", row.ChildText(`td`))
				}
			})
		})
	})

	c.Visit("https://www.gov.uk/guidance/dvla-coronavirus-covid-19-update")
}
