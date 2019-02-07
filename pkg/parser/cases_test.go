package parser

type setTest struct {
	input    string
	expected error
}

type stringTest struct {
	expected string
}

var setTestCases = []setTest{
	{
		input:    "http://foo.bar",
		expected: nil,
	},
	{
		input:    "http://bibidi.babido.boo",
		expected: nil,
	},
}

var stringTestCases = []stringTest{
	{
		expected: "",
	},
}
