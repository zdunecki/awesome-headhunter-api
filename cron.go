package main

import (
	"gopkg.in/robfig/cron.v3"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	c := cron.New()
	forever := make(chan bool)

	_, _ = c.AddFunc("@every 1m", facebookCrawl)

	c.Start()
	<-forever
}

func facebookCrawl() {
	log.Printf("%s", "Start Facebook crawling")
	resp, err := http.Get("http://localhost:3000/api/facebook/crawler")

	if err != nil {
		log.Printf("%s", "Something went wrong")

		panic(err)
	}

	defer resp.Body.Close()

	log.Printf("%s", "Read body stream")
	b, err := ioutil.ReadAll(resp.Body)

	log.Printf("%s", b)
}
