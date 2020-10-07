package search

import (
	"container/list"
	"fmt"
	"net/url"
)

// BFS is Breadth-first search implementation
func BFS(rootURL string) error { // O(avg_branching_factor**(depth + 1) )
	rURL, err := url.Parse(rootURL)
	if err != nil {
		return err
	}
	queue := list.New()
	marked := make(map[string]bool)
	marked[rootURL] = true
	queue.PushBack(rootURL) // Enqueue
	for queue.Len() > 0 {
		e := queue.Front()
		if e == nil {
			return nil
		}
		queue.Remove(e)      // Dequeue
		fmt.Println(e.Value) // Visit
		if strValue, ok := e.Value.(string); ok {
			links, err := getLinks(rootURL, strValue)
			if err != nil {
				return err
			}
			for _, link := range links {
				l, err := url.Parse(link)
				if err != nil {
					// decide to continue scanning for URLs and ignore the link found that does not match an URL format
					continue
				}
				_, found := marked[link]
				// Enqueue the new url if it is not marked already (duplicated) and its domain is the same as the rootURL domain
				if !found && rURL.Hostname() == l.Hostname() {
					marked[link] = true
					queue.PushBack(link) // Enqueue
				}
			}
		}
	}
	return nil
}
