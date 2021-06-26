package engine

import (
	"learn-golang/crawler/scheduler"
	"learn-golang/crawler/types"
	"log"
)

type ConcurrentEngine struct {
	Scheduler   scheduler.Scheduler
	WorkerCount int
	ItemChan chan types.Item
	RequestProcessor Processor
}

type Processor func (r types.Request) (types.ParseResult, error)

// 并发版
func (c *ConcurrentEngine) Run(seeds ...types.Request) {
	out := make(chan types.ParseResult) //数据结果通道
	c.Scheduler.Run()                   //执行任务调度,

	for i := 0; i < c.WorkerCount; i++ {
		//创建多个work, work从调度器的request通道获得数据, 向result通道发送信息
		c.createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
		//向请求通道提交request
		c.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		//result通道有数据过来了
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %+v", itemCount, item)
			itemCount++
			c.ItemChan <- item
		}

		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			c.Scheduler.Submit(request)
		}
	}
}

var visitedUrl = make(map[string]bool)

// URL deduplicate
func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}


func (c ConcurrentEngine) createWorker(in chan types.Request, out chan types.ParseResult, ready scheduler.ReadyNotifier) {
	go func() {
		for {
			// Tell scheduler I am ready
			ready.WorkerReady(in)
			request := <-in
			result, err := c.RequestProcessor(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
