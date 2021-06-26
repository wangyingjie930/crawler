package types

type Request struct {
	Url       string
	//ParseFunc ParseFunc
	Parser Parser
}

type Parser interface {
	Serialize() (name string, args interface{})
	Parse(content []byte, url string) ParseResult
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

type FuncParser struct {
	parser ParseFunc
	name   string
}

func (f *FuncParser) Parse(contents []byte, url string) ParseResult {
	return f.parser(contents, url)
}

func (f *FuncParser) Serialize() (name string, args interface{}) {
	return f.name, nil
}

// 工厂函数创建FuncParser
func NewFuncParser(p ParseFunc, name string) *FuncParser {
	return &FuncParser{
		parser: p,
		name:   name,
	}
}