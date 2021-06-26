package client

import (
	"learn-golang/crawler-distribute/worker"
	"learn-golang/crawler/engine"
	"learn-golang/crawler/types"
	"net/rpc"
)


func CreateProcessor(clientChan chan *rpc.Client) engine.Processor {

	return func(r types.Request) (types.ParseResult, error) {
		request := worker.SerializeRequest(r)
		var result worker.ParseResult
		client := <- clientChan
		err := client.Call("CrawService.Process", request, &result)
		if err != nil {
			return types.ParseResult{}, err
		}
		return worker.DeserializeResult(result), nil
	}
}
