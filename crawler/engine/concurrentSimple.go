package engine

import (
	"fmt"
	"learn-golang/crawler/scheduler"
	"learn-golang/crawler/types"
	"log"
)

type ConcurrentSimpleEngine struct {
	Scheduler   scheduler.Scheduler
	WorkerCount int
}

// 并发版
func (e *ConcurrentSimpleEngine) Run(seeds ...types.Request) {
	e.Scheduler.Run()

	out := make(chan types.ParseResult)
	for i := 1; i <= e.WorkerCount; i ++ {
		e.createWork(e.Scheduler.WorkerChan(), out)
	}
	for _, seed := range seeds {
		e.Scheduler.Submit(seed)
	}

	for {
		result := <- out
		log.Printf("%+v", result)
		for _, item := range result.Items {
			fmt.Printf("Got item %+v \n", item)
		}

		for _, request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func (e *ConcurrentSimpleEngine) createWork(in chan types.Request, out chan types.ParseResult)  {
	go func() {
		for {
			request := <- in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			out <- parseResult
		}
	}()
}
