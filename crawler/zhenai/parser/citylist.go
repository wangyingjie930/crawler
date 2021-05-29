package parser

import (
	"learn-golang/crawler/types"
	"regexp"
)

var CityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

// 获取城市列表
func ParseCityList(contents []byte, _ string) types.ParseResult {
	matches := CityListRe.FindAllSubmatch(contents, -1)
	//fmt.Printf("content: %s", contents)
	result := types.ParseResult{}
	count := 0
	for _, m := range matches {
		count ++
		result.Requests = append(
			result.Requests,
			types.Request{Url: string(m[1]), ParseFunc: ParseCityUserList})
		if count > 10 {
			break
		}
	}
	return result
}
