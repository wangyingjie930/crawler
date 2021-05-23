package main

import (
	"learn-golang/crawler/engine"
	"learn-golang/crawler/scheduler"
	"learn-golang/crawler/types"
	"learn-golang/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		//Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
