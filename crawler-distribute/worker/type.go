package worker

import (
	"errors"
	"learn-golang/crawler/types"
	parser2 "learn-golang/crawler/zhenai/parser"
	"log"
)

type Request struct {
	Url string
	Parser SerializedParser
}

type SerializedParser struct {
	Name string
	Args interface{}
}

type ParseResult struct {
	Requests []Request
	Items    []types.Item
}

func DeserializeRequest(r Request) (types.Request, error)  {
	parser, err := deserializeParser(r.Parser)
	if err != nil {
		return types.Request{}, err
	}
	return types.Request{
		Url: r.Url,
		Parser: parser,
	}, nil
}

func DeserializeResult(r ParseResult) types.ParseResult {
	result := types.ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		engineReq, err := DeserializeRequest(req)
		if err != nil {
			//其中一个Request加错了就不加Request
			log.Printf("error deserializing "+"request: %v", err)
			continue
		}
		result.Requests = append(result.Requests, engineReq)
	}
	return result
}

func SerializeResult(r types.ParseResult) ParseResult {
	result := ParseResult{
		Items: r.Items,
	}
	for _, req := range r.Requests {
		result.Requests = append(result.Requests, SerializeRequest(req))
	}
	return result
}

func SerializeRequest(r types.Request) Request  {
	name, args := r.Parser.Serialize()
	return Request{
		Url: r.Url,
		Parser: SerializedParser{
			Name: name,
			Args: args,
		},
	}
}

func deserializeParser(parser SerializedParser) (types.Parser, error) {
	switch parser.Name {
	case "ParseCityUserList":
		return types.NewFuncParser(parser2.ParseCityUserList, "ParseCityUserList"), nil
	case "ParseCityList":
		return types.NewFuncParser(parser2.ParseCityList, "ParseCityList"), nil
	case "ProfileParser":
		return parser2.NewProfileParser(parser.Args.(string)), nil
	default:
		return nil, errors.New("找不到该调用函数")
	}
}