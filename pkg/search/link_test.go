package search

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.Nil(t, err)
	assert.ObjectsAreEqual(links, expected)
}

func TestRemoveDuplicates(t *testing.T) {
	for _, tt := range removeDuplicatesTestCases {
		result := removeDuplicates(tt.input)
		assert.ObjectsAreEqual(result, tt.expected)
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range removeDuplicatesTestCases {
			removeDuplicates(test.input)
		}
	}
}
