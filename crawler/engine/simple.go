package engine

import (
	"learn-golang/crawler/types"
	"log"
)

type SimpleEngine struct{}

// 串行
func (e SimpleEngine) Run(seeds ...types.Request) {
	var requests []types.Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := Worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		//out := persist.ItemServer()
		for _, item := range parseResult.Items {
			log.Printf("Got item %+v", item)
			//out <- item
		}
	}
}
