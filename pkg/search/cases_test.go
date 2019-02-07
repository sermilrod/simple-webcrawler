package search

type removeDuplicatesTest struct {
	input       []string
	expected    []string
	description string
}

var removeDuplicatesTestCases = []removeDuplicatesTest{
	{
		input: []string{
			"http://foo.bar",
			"http://foo.bar",
			"http://foo.bar/bibidi",
			"http://foo.bar/babidi",
			"http://foo.bar/boo",
			"http://foo.bar/bibidi",
			"http://foo.bar/babidi",
			"http://foo.bar/boo",
		},
		expected: []string{
			"http://foo.bar",
			"http://foo.bar/bibidi",
			"http://foo.bar/babidi",
			"http://foo.bar/boo",
		},
		description: "Slice with duplicates",
	},
	{
		input:       []string{},
		expected:    []string{},
		description: "Empty slice",
	},
	{
		input:       []string{"http://foo.bar"},
		expected:    []string{"http://foo.bar"},
		description: "Slice without duplicates",
	},
}
