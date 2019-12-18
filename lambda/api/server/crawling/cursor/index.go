package cursor

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const DefaultCursor = "https://facebook.com/MONOstudio-393020864107976/"

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

func Handler(w http.ResponseWriter, r *http.Request) {
	path := bfsCursor()

	b, err := json.Marshal(path)

	if err != nil {
		panic(err)
	}

	if _, err := w.Write(b); err != nil {
		panic(err)
	}
}

func bfsCursor() []string {
	start := DefaultCursor

	graph, nodes := writeGraph()
	path := bfs(start, graph)

	var finalPath []string

	for _, p := range path {
		if contains(nodes, p) {
			continue
		}
		finalPath = append(finalPath, p)
	}

	return finalPath
}

func writeGraph() (g map[string][]string, n []string) {
	var graph = make(map[string][]string)

	// fill bfs graph
	var graphNodes []string
	nodes := client.Keys("crawling|graph|*").Val()

	for _, node := range nodes {
		leafs := client.SMembers(node).Val()

		graphNode := strings.Split(node, "|")[2]
		graph[graphNode] = leafs
		graphNodes = append(graphNodes, graphNode)
	}

	return graph, graphNodes
}

func saveGraph(cursor string, data []string) {
	linksI := make([]interface{}, len(data))
	for i, v := range data {
		linksI[i] = v
	}

	client.SAdd("crawling|graph|"+cursor, linksI...)
}

func bfs(start string, graph map[string][]string) []string {
	var path []string

	queue := []string{start}

	// bfs
	for {
		if len(queue) <= 0 {
			break
		}
		node := queue[0]
		queue = queue[1:]

		if contains(path, node) {
			continue
		}

		var neighbours []string

		path = append(path, node)
		neighbours = graph[node]

		for _, neighbour := range neighbours {
			if contains(path, neighbour) {
				continue
			}
			queue = append(queue, neighbour)
		}
	}

	return path
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
