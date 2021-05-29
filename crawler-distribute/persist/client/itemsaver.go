package client

import (
	"learn-golang/crawler-distribute/rpcsupport"
	"learn-golang/crawler/types"
	"log"
)

func ItemSaver(host string) chan types.Item {
	client := rpcsupport.NewClient(host)
	out := make(chan types.Item)
	go func() {
		itemCount := 0
		for  {
			item := <-out
			log.Printf("Got item #%d: %+v", itemCount, item)
			itemCount++

			var result string
			err := client.Call("ItemSaverService.Save", item, &result)
			if err != nil {
				log.Printf("save error %v", err)
			}
		}
	}()
	return out
}
