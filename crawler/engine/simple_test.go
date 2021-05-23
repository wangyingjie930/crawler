package engine

import (
	"learn-golang/crawler/types"
	"learn-golang/crawler/zhenai/parser"
	"testing"
)

func TestSimpleEngine_Run(t *testing.T) {
	SimpleEngine{}.Run(types.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}

func TestSimpleSample_Run(t *testing.T) {
	SimpleSample{}.Run(types.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
