package main

import (
	"bytes"
	"encoding/json"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const SuggestedPagesSelector = "._5ay5[data-id='10']"
const ProfileLinkSelector = "a._8o._8t._ohe"

func Handler(w http.ResponseWriter, r *http.Request) {
	cursor := getCrawlingCursor()

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML(SuggestedPagesSelector, onSuggestedPages)

	if err := c.Visit(string(cursor)); err != nil {
		panic(err)
	}

	w.Write([]byte("ok"))
}

func getCrawlingCursor() string {
	response, _ := http.Get(
		"http://localhost:3000/api/server/crawling/cursor",
	)

	defer response.Body.Close()

	cursor, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	return string(cursor)
}

func onSuggestedPages(e *colly.HTMLElement) {
	var links []string

	e.ForEach(ProfileLinkSelector, func(i int, element *colly.HTMLElement) {
		href := element.Attr("href")

		re := regexp.MustCompile(`^[^?]+`)
		link := re.FindString(href)

		log.Printf("%s", link)

		links = append(links, link)
	})

	jsonLinks, err := json.Marshal(links)

	if err != nil {
		panic(err)
	}

	response, _ := http.Post(
		"http://localhost:3000/api/server/crawling",
		"application/json",
		bytes.NewBuffer(jsonLinks),
	)

	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
}
