package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"learn-golang/crawler/types"
	"log"
)

func ItemServer(index string) chan types.Item {
	client, err := elastic.NewClient(
		elastic.SetSniff(false), elastic.SetURL("http://192.168.205.10:9200"))
	if err != nil {
		panic(err)
	}

	out := make(chan types.Item)
	go func() {
		itemCount := 0
		for  {
			item := <-out
			log.Printf("Got item #%d: %+v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				log.Print("save error")
			}
		}
	}()
	return out
}

func Save(client *elastic.Client, index string, item types.Item) error {
	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	resp, err := indexService.Do(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("%+v", resp)
	return nil
}
