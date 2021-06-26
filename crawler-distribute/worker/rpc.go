package worker

import (
	"learn-golang/crawler/engine"
)

type CrawService struct {}

func (CrawService) Process (req Request, result *interface{}) error {
	engineReq, err := DeserializeRequest(req)
	if err != nil {
		return err
	}
	engineResult, err := engine.Worker(engineReq)
	if err != nil {
		return err
	}
	*result = SerializeResult(engineResult)
	return nil
}
