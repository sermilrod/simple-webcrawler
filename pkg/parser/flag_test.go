package parser

import (
	"testing"
)

func TestURLListString(t *testing.T) {
	urlList := make(URLList, 0)
	for _, test := range stringTestCases {
		if result := urlList.String(); result != test.expected {
			t.Errorf("URLListString test: expected [%s], got [%s]", test.expected, result)
		}
	}
}

func TestURLListSet(t *testing.T) {
	urlList := make(URLList, 0)
	for _, test := range setTestCases {
		if result := urlList.Set(test.input); result != test.expected {
			t.Errorf("URLListSet test: expected [%s], got [%s]", test.expected, result)
		}
	}
	if len(urlList) != len(setTestCases) {
		t.Errorf("URLListSet test: expected [%d], got [%d]", len(setTestCases), len(urlList))
	}
}

func BenchmarkSet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlList := make(URLList, 0)
		for _, test := range setTestCases {
			urlList.Set(test.input)
		}
	}
}
