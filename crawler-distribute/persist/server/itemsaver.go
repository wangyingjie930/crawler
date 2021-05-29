package main

import (
	"flag"
	"fmt"
	"github.com/olivere/elastic"
	"learn-golang/crawler-distribute/config"
	"learn-golang/crawler-distribute/persist"
	"learn-golang/crawler-distribute/rpcsupport"
	"log"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main()  {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
	}
	// 出错强制退出
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host string, index string) error {
	client, err := elastic.NewClient(
		elastic.SetSniff(false), elastic.SetURL("http://192.168.205.10:9200"))
	if err != nil {
		return err
	}
	err = rpcsupport.ServerRpc(host, persist.ItemSaverService{
		Client: client,
		Index: index,
	})
	if err != nil {
		return err
	}
	return nil
}
