package crawling

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Data struct {
	Cursor     string   `json:"cursor"`
	Links      []string `json:"links"`
	Instagram  string   `json:"instagram"`
	Categories []string `json:"categories"`
}

var client *redis.Client

func init() {
	redisUrl, _ := url.Parse(os.Getenv("REDIS_URL"))
	redisPassword, _ := redisUrl.User.Password()
	redisDB := 0

	redisOptions := redis.Options{
		Addr:     redisUrl.Host,
		Password: redisPassword,
		DB:       redisDB,
	}

	client = redis.NewClient(&redisOptions)
}

func Abc() int {
	return 5
}

var ProfileCategories = []string{"interior-designer", "architectural-designer"}

func Handler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var data []Data

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		panic(err)
	}

	saveGraph(data)

	w.Write([]byte("ok"))
}

// {
// 	 "crawling:graph:<cursor1>": links
//	 "crawling:graph:<cursorN+1>": links
// }
func saveGraph(crawlerData []Data) {
	for _, data := range crawlerData {
		profile := struct {
			Instagram  string   `json:"instagram"`
			Categories []string `json:"categories"`
		}{
			data.Instagram,
			data.Categories,
		}

		jsonProfile, e := json.Marshal(profile)
		if e != nil {
			panic(e)
		}

		client.Set("crawling|profile|"+data.Cursor, jsonProfile, -1)

		canSaveGraph := false
		for _, c := range profile.Categories {
			if contains(ProfileCategories, c) {
				canSaveGraph = true
				break
			}
		}

		linksI := make([]interface{}, len(data.Links))
		for i, v := range data.Links {
			linksI[i] = v
		}

		if canSaveGraph {
			client.SAdd("crawling|graph|"+data.Cursor, linksI...)
		} else { // remove not satisfied node neighbours
			nodes := client.Keys("crawling|graph|*").Val()

			for _, node := range nodes {
				client.SRem(node, data.Cursor)
			}
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
