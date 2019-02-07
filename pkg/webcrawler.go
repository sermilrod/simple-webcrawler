package webcrawler

import (
	"flag"

	"github.com/labstack/gommon/log"
	"github.com/sermilrod/simple-webcrawler/pkg/parser"
	"github.com/sermilrod/simple-webcrawler/pkg/search"
)

var urls parser.URLList

func Start() {
	flag.Var(&urls, "url", "URL to crawl")
	flag.Parse()

	for _, url := range urls {
		// Traversing the links of a website is equivalent to traverse a tree/graph.
		// The URL is the root node and each link is a neighbour.
		// We can resolve this problem by applying Breadth-first search (BFS).
		// BFS will explore all of the neighbour links at the present depth before moving on to the links at the next depth level.
		if err := search.BFS(url); err != nil {
			log.Warnf("Unable to crawl %s. Error: %s", url, err.Error())
		}
	}
}
