package utils

import (
	"encoding/json"
	"github.com/wudian/GoEx"
	"github.com/wudian/market_data/models"
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

func Struct2JsonBytes(structt interface{}) (jsonBytes []byte, err error) {
	data, err := json.Marshal(structt)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GoexTicker2Ticker(goexTick *goex.Ticker, api string) *models.Ticker {
	if nil == goexTick {
		goexTick = goex.NewTicker()
	}
	ticker := &models.Ticker{
		Api:    api,
		Symbol: goexTick.Pair.ToSymbol("-"),
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

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	} else if valueSlice, ok := value.(map[string]float64); ok {
		newSlice := make(map[string]float64)
		for k, v := range valueSlice {
			newSlice[k] = v
		}

		return newSlice
	}

	return value
}
