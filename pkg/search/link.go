package search

import (
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getLinks(rootURL, url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		return nil, err
	}
	links, err := extractLinks(rootURL, resp)
	if err != nil {
		defer resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return links, nil
}

func extractLinks(rootURL string, resp *http.Response) ([]string, error) {
	links := make([]string, 0)
	u, err := url.Parse(rootURL)
	if err != nil {
		return nil, err
	}
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return removeDuplicates(links), nil
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link, err := resp.Request.URL.Parse(attr.Val)
						if err != nil || !strings.Contains(link.Hostname(), u.Hostname()) {
							continue
						}
						links = append(links, link.String())
					}
				}
			}
		}
	}
}

func removeDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := []string{}
	for v := range elements {
		if encountered[elements[v]] == false {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}
	return result
}
