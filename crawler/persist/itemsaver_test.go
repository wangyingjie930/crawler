package persist

import (
	"learn-golang/crawler/zhenai/model"
	"testing"
)

func TestSave(t *testing.T) {
	item := model.Profile{
		Name:       "惠儿",
		Age:        50,
		Height:     156,
		Weight:     0,
		Income:     "3000元以下",
		Gender:     "女",
		Xinzuo:     "魔羯座",
		Marriage:   "离异",
		Education:  "高中及以下",
		Occupation: "销售总监",
		Hokou:      "四川阿坝",
		House:      "租房",
		Car:        "未购车",
	}
	save(item)
}
