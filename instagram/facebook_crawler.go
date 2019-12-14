package main

import (
	"github.com/go-redis/redis"
	"github.com/gocolly/colly"
	"log"
	"regexp"
)

const SuggestedPagesSelector = "._5ay5[data-id='10']"
const ProfileLinkSelector = "a._8o._8t._ohe"

func GetFacebookCrawlerCursor() string {
	return "https://pl-pl.facebook.com/pages/category/Interior-Design-Studio/MONOstudio-393020864107976/"
}

func main() {
	var client *redis.Client

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	var links []string

	c.OnHTML(SuggestedPagesSelector, func(e *colly.HTMLElement) {
		e.ForEach(ProfileLinkSelector, func(i int, element *colly.HTMLElement) {
			href := element.Attr("href")

			re := regexp.MustCompile(`^[^?]+`)
			link := re.FindString(href)

			log.Printf("%s", link)

			links = append(links, link)
		})

		client.Publish("crawling", links)
	})

	if err := c.Visit(GetFacebookCrawlerCursor()); err != nil {
		panic(err)
	}
}
