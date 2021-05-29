package types

type Request struct {
	Url       string
	ParseFunc ParseFunc
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

func NilParser([]byte) ParseResult {
	return ParseResult{}
}

type Item struct {
	Url     string
	Type    string
	Id      string
	Payload interface{}
}

type ParseFunc func(content []byte, url string) ParseResult
