package main

import (
	"github.com/go-redis/redis"
	"log"
)

func main() {
	var client *redis.Client

	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pubsub := client.PSubscribe("scraping", "crawling")

	defer pubsub.Close()

	go func() {
		msg := pubsub.Channel()
		for {
			ch := <-msg

			switch ch.Channel {
			case "scraping":
				{
					log.Printf("%s", ch.Payload)
				}
			case "crawling":
				{
					log.Printf("%s", ch.Payload)
				}
			}
		}

	}()

	infinite := make(chan bool)

	<-infinite
}
