package main

import (
	"flag"
	"fmt"
	"learn-golang/crawler-distribute/rpcsupport"
	"learn-golang/crawler-distribute/worker"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	log.Fatal(rpcsupport.ServerRpc(fmt.Sprintf(":%d", *port), worker.CrawService{}))
}
