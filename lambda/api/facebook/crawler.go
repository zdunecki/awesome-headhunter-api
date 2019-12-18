package main

import (
	"bytes"
	"encoding/json"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Data struct {
	Cursor     string   `json:"cursor"`
	Links      []string `json:"links"`
	Instagram  string   `json:"instagram"`
	Categories []string `json:"categories"`
}

const SuggestedPagesSelector = "._5ay5[data-id='10']"
const ProfileLinkSelector = "a._8o._8t._ohe"
const ProfileInstagramSelector = "._4bl9 > a[href*='instagram']"
const ProfileCategorySelector = "._4bl9 > a[href*='/pages/category/']"

var HeadhunterURL = os.Getenv("HEADHUNTER_URL")
var CatchYCursors = 2

//TODO: make it better, we should keep proxies as map for better crawler results
// for example map of geolocation and proxy url and make shake with roundrobin algorithm on colly
var CrawlerProxies []string

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := json.Unmarshal([]byte(os.Getenv("CRAWLER_PROXIES")), &CrawlerProxies); err != nil {
		panic(err)
	}

	cursors := getCrawlingCursor()

	if len(cursors) > 1 {
		cursors = cursors[0:CatchYCursors]
	}

	crawlerData := make([]Data, len(cursors))

	for i, cursor := range cursors {
		crawlerData[i] = Data{
			Cursor: cursor,
		}

		scrapePage(cursor, func(links []string) {
			crawlerData[i].Links = links
		})
		scrapeAbout(
			cursor,
			func(category string) {
				crawlerData[i].Categories = append(crawlerData[i].Categories, category)
			},
			func(instagram string) {
				crawlerData[i].Instagram = instagram
			})
	}

	if err := sendCrawlingData(crawlerData); err != nil {
		w.WriteHeader(500)
		panic(err)
		return
	}

	b, _ := json.Marshal(cursors)
	w.Write(b)

}

func scrapePage(cursor string, cb func([]string)) {
	c := colly.NewCollector()
	if err := c.SetProxy(CrawlerProxies[0]); err != nil {
		log.Fatal(err)
	}

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML(
		SuggestedPagesSelector,
		func(e *colly.HTMLElement) {
			if cursor == "" {
				return
			}
			var links []string

			e.ForEach(ProfileLinkSelector, func(i int, element *colly.HTMLElement) {
				href := element.Attr("href")

				re := regexp.MustCompile(`^[^?]+`)
				link := re.FindString(href)

				links = append(links, link)
			})

			cb(links)
		},
	)

	if err := c.Visit(cursor); err != nil {
		panic(err)
	}

	c.Wait()
}

func scrapeAbout(cursor string, cb func(string), instagramCb func(string)) {
	onProfileCategory := func(e *colly.HTMLElement) {
		categoryLink := e.Attr("href")
		category := strings.Replace(
			strings.Split(categoryLink, "/category")[1],
			"/",
			"",
			-1,
		)
		cb(category)
	}

	c := colly.NewCollector()
	if err := c.SetProxy(CrawlerProxies[0]); err != nil {
		log.Fatal(err)
	}

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	c.OnHTML(
		ProfileInstagramSelector,
		func(e *colly.HTMLElement) {
			instagramCb(e.Attr("href"))
		},
	)
	c.OnHTML(
		ProfileCategorySelector,
		onProfileCategory,
	)

	if err := c.Visit(cursor + "about"); err != nil {
		panic(err)
	}

	c.Wait()
}

func getCrawlingCursor() []string {
	response, _ := http.Get(
		HeadhunterURL + "/api/server/crawling/cursor",
	)

	defer response.Body.Close()

	cursor, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	var links []string

	if err := json.Unmarshal(cursor, &links); err != nil {
		panic(err)
	}

	return links
}

func sendCrawlingData(crawlerData []Data) error {
	jsonData, err := json.Marshal(crawlerData)

	response, _ := http.Post(
		HeadhunterURL+"/api/server/crawling",
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	defer response.Body.Close()

	_, err = ioutil.ReadAll(response.Body)

	return err
}
