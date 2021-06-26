package main

import (
	"flag"
	client2 "learn-golang/crawler-distribute/persist/client"
	"learn-golang/crawler-distribute/rpcsupport"
	"learn-golang/crawler-distribute/worker/client"
	"learn-golang/crawler/engine"
	"learn-golang/crawler/scheduler"
	"learn-golang/crawler/types"
	"learn-golang/crawler/zhenai/parser"
	"net/rpc"
	"strings"
)

var (
	// go run main.go --itemsaver_host=":1234" --worker_hosts=":9000,:9001"
	itemSaverHost = flag.String(
		"itemsaver_host", "", "itemsaver host")
	workerHosts = flag.String("worker_hosts", "", "worker hosts (comma separated)")
)

func main() {
	flag.Parse()
	itemSaver := client2.ItemSaver(*itemSaverHost)

	pool := createClientPool(strings.Split(*workerHosts, ", "))
	processor := client.CreateProcessor(pool)

	eng := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan: itemSaver,
		RequestProcessor: processor,
	}

	eng.Run(types.Request{
		Url: "http://www.zhenai.com/zhenghun",
		Parser: types.NewFuncParser(parser.ParseCityList, "ParseCityList"),
	})
}

func createClientPool(hosts []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, host := range hosts{
		clients = append(clients, rpcsupport.NewClient(host))
	}
	var clientChan chan *rpc.Client
	go func() {
		for {
			for _, v := range clients{
				clientChan <- v
			}
		}
	}()
	return clientChan
}
