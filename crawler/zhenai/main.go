package main

import (
	"learn-golang/crawler-distribute/persist/client"
	"learn-golang/crawler/engine"
	"learn-golang/crawler/scheduler"
	"learn-golang/crawler/types"
	"learn-golang/crawler/zhenai/parser"
)

func main() {
	//itemChan := persist.ItemServer("dating_profile")
	itemChan := client.ItemSaver(":1234")

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		//Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
		ItemChan: itemChan,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
