package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLListString(t *testing.T) {
	urlList := make(URLList, 0)
	for _, tt := range stringTestCases {
		result := urlList.String()
		assert.Equal(t, tt.expected, result)
	}
}

func TestURLListSet(t *testing.T) {
	urlList := make(URLList, 0)
	for _, tt := range setTestCases {
		result := urlList.Set(tt.input)
		assert.Equal(t, tt.expected, result)
	}
	assert.Equal(t, len(urlList), len(setTestCases))
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlList := make(URLList, 0)
		for _, test := range setTestCases {
			urlList.Set(test.input)
		}
	}
}
