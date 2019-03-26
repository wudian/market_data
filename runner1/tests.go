package main

import (
	"encoding/json"
	"fmt"
	"github.com/huobiapi/REST-GO-demo/services"
	"github.com/okcoin-okex/open-api-v3-sdk/okex-go-sdk-api"
)

func test2()  {
	c := okex.NewTestClient()
	ac, err := c.GetSpotInstrumentTicker("LTC-USDT")
	if err == nil {
		jstr, _ := okex.Struct2JsonString(ac)
		println(jstr)
	}
}
func test()  {
	ret := services.GetSymbols()
	jsonStr, err := json.Marshal(ret)
	if err == nil {
		fmt.Println(string(jsonStr))
	}
}

func test3()  {
	fmt.Println("4444")
}