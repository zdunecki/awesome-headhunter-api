package main

import (
	"context"
	"encoding/json"
	"github.com/chromedp/chromedp"
	"github.com/go-redis/redis"
	"log"
)

type InstagramSharedData struct {
	EntryData struct {
		ProfilePage []struct {
			Graphql struct {
				User struct {
					Username             string `json:"username"`
					Biography            string `json:"biography"`
					BusinessCategoryName string `json:"business_category_name"`
					EdgeFollow           struct {
						Count int `json:"count"`
					} `json:"edge_follow"`
				} `json:"user"`
			} `json:"graphql"`
		} `json:"ProfilePage"`
	} `json:"entry_data"`
}

func main() {
	var client *redis.Client

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var ig InstagramSharedData

	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.instagram.com/saints/"),
		chromedp.Evaluate(`window._sharedData;`, &ig),
	)
	if err != nil {
		log.Fatal(err)
	}

	queueData, err := json.Marshal(ig)
	if err != nil {
		log.Fatal(err)
	}

	client.Publish("scraping", queueData)
	log.Print("send data to queue")
}
