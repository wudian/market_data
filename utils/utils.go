package utils

import (
	"encoding/json"
	"github.com/nntaoli-project/GoEx"
	"github.com/wudian/wx/models"
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

func GoexTicker2Ticker(goexTick *goex.Ticker) *models.Ticker {
	ticker := &models.Ticker{
		Symbol: goexTick.Pair.ToSymbol("_"),
		Last: goexTick.Last,
		Buy: goexTick.Buy,
		Sell: goexTick.Sell,
		High: goexTick.High,
		Low: goexTick.Low,
		Vol: goexTick.Vol,
		Date: goexTick.Date,
	}
	return ticker
}