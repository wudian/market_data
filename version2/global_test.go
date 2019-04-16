package main

import (
	"github.com/wudian/GoEx"
	"github.com/stretchr/testify/assert"
	"testing"
)

func testSymbol(t *testing.T) {
	global := GlobalInstance()
	for _, api := range global.apiNames {
		for _, symbol := range global.vecSymbols {
			t.Log(symbol + "->"+api+"->"+goex.NewCurrencyPair2(symbol).String())
		}
	}

}

func TestApi(t *testing.T)  {
	global := GlobalInstance()
	goexTicker, err := global.apis["binance"].GetTicker(goex.BTC_USDT)
	jsonStr, err := Struct2JsonString(goexTicker)
	if err == nil {
		t.Log(string(jsonStr))
	}

	ticker := GoexTicker2Ticker(goexTicker)
	jsonStr, err = Struct2JsonString(ticker)
	if err == nil {
		t.Log(string(jsonStr))
	}
}

func testKafka(t *testing.T)  {
	syncProducer()
	assert.Equal(t, "123", "123")
}