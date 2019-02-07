package search

import (
	"container/list"
	"fmt"
)

// BFS is Breadth-first search implementation
func BFS(rootURL string) error { // O(avg_branching_factor**(depth + 1) )
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
				if _, ok := marked[link]; !ok {
					marked[link] = true
					queue.PushBack(link) // Enqueue
				}
			}
		}
	}
	return nil
}
