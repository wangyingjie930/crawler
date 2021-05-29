package persist

import (
	"github.com/olivere/elastic"
	"learn-golang/crawler/persist"
	"learn-golang/crawler/types"
	"log"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (i ItemSaverService) Save (item types.Item, result *interface{}) error {
	err := persist.Save(i.Client, i.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return nil
}
