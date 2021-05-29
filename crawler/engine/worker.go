package engine

import (
	"learn-golang/crawler/fetcher"
	"learn-golang/crawler/types"
	"log"
)

func worker(r types.Request) (types.ParseResult, error) {
	// log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v\n", r.Url, err)
		return types.ParseResult{}, err
	}
	return r.ParseFunc(body, r.Url), nil
}
