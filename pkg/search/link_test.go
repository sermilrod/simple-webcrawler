package search

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetLinks(t *testing.T) {
	// Start a local test HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		htmlPage := []byte(`
<html>
	<head>
		<title>
			Title of the document
		</title>
	</head>
	<body>
		Some content
		<a href="http://127.0.0.1/bibidi"></a>
	</body>
</html>
`)
		rw.Write(htmlPage)
	}))
	defer server.Close()

	links, err := getLinks(server.URL, server.URL)
	expected := []string{"http://127.0.0.1/bibidi"}
	if err != nil {
		t.Errorf("getLinks test: expected err to be nil, got [%s]", err.Error())
	}
	if !reflect.DeepEqual(links, expected) {
		t.Errorf("getLinks test: expected links %s, got %s", expected, links)
	}
}

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range removeDuplicatesTestCases {
		result := removeDuplicates(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("RemoveDuplicates test: %s, expected [%s], got [%s]", test.description, test.expected, result)
		}
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range removeDuplicatesTestCases {
			removeDuplicates(test.input)
		}
	}
}
