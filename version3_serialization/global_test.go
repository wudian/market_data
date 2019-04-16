package main

import (
	"github.com/wudian/GoEx"
	"github.com/stretchr/testify/assert"
	"github.com/wudian/wx/config"
	"github.com/wudian/wx/kafka"
	"github.com/wudian/wx/utils"
	"testing"
)

func testSymbol(t *testing.T) {
	global := config.GlobalInstance()
	for _, api := range global.ApiNames {
		for _, symbol := range global.VecSymbols {
			t.Log(symbol + "->"+api+"->"+goex.NewCurrencyPair2(symbol).String())
		}
	}

}

func TestApi(t *testing.T)  {
	global := config.GlobalInstance()
	goexTicker, err := global.Apis["binance"].GetTicker(goex.BTC_USDT)
	jsonStr, err := utils.Struct2JsonString(goexTicker)
	if err == nil {
		t.Log(string(jsonStr))
	}

	ticker := utils.GoexTicker2Ticker(goexTicker)
	jsonStr, err = utils.Struct2JsonString(ticker)
	if err == nil {
		t.Log(string(jsonStr))
	}
}

func testKafka(t *testing.T)  {
	kafka.SyncProducer()
	assert.Equal(t, "123", "123")
}