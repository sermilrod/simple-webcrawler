package parser

// Type that represents the list of URLs to crawl
type URLList []string

// String satisfies the flag.Value interface.
func (urlsList *URLList) String() string {
	return ""
}

// Set is called once, in command line order, for each flag present.
// It is required to satisfy the flag.Value interface.
// Everytime there is a new flag with the same name we append its value.
func (urlsList *URLList) Set(value string) error {
	*urlsList = append(*urlsList, value)
	return nil
}
