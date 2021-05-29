package main

import (
	"learn-golang/crawler-distribute/config"
	"learn-golang/crawler-distribute/rpcsupport"
	"learn-golang/crawler/types"
	"learn-golang/crawler/zhenai/model"
	"testing"
	"time"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"
	// start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second * 2)
	// start ItemSaverClient
	client := rpcsupport.NewClient(host)
	// Call save
	item := types.Item{
		Url:  "https://album.zhenai.com/u/1221657259",
		Type: "zhenai",
		Id:   "1089023121",
		Payload: model.Profile{
			Name:          "芜湖小啊妹",
			Gender:        "女",
		},
	}
	result := ""
	err := client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err: %s", result, err)
	}
}
