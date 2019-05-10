package main

import (
	"encoding/json"
	"github.com/wudian/GoEx"
)

/*
 struct convert json string
*/
func Struct2JsonString(structt interface{}) (jsonString string, err error) {
	data, err := json.Marshal(structt)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func GoexTicker2Ticker(goexTick *goex.Ticker) *Ticker {
	ticker := &Ticker{
		Symbol: goexTick.Pair.ToSymbol("_"),
		Last:   goexTick.Last,
		Buy:    goexTick.Buy,
		Sell:   goexTick.Sell,
		High:   goexTick.High,
		Low:    goexTick.Low,
		Vol:    goexTick.Vol,
		Date:   goexTick.Date,
	}
	return ticker
}
