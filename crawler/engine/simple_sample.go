package engine

import (
	"fmt"
	"learn-golang/crawler/fetcher"
	"learn-golang/crawler/types"
	"log"
)

type SimpleSample struct {
	
}

//合并数组
//获取数组第一个元素, pop

func (s SimpleSample) Run (seed ...types.Request) {
	var requests []types.Request
	requests = append(requests, seed...)

	for len(requests) > 0 {
		request := requests[0]
		requests = requests[1:]

		res, err := fetcher.Fetch(request.Url)
		if err != nil {
			log.Print("error request!!!")
			//continue
		}
		//第一次返回: requests: [{url: 城市链接, parseFunc: ParseCityUserList}], items: []
		//第二次返回: requests: [{url: 用户链接, parseFunc: ParseProfile}], items: []
		//第三次返回: requests: [], items: [......]
		parseResult := request.Parser.Parse(res, request.Url)
		log.Printf("解析结果: %+v", parseResult)
		for _, item := range parseResult.Items {
			fmt.Printf("print item %+v \n", item)
		}
		requests = append(requests, parseResult.Requests...)
	}
}