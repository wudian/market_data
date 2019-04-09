package main

import (
	"github.com/nntaoli-project/GoEx"
	"github.com/stretchr/testify/assert"
	"github.com/wudian/wx/config"
	"github.com/wudian/wx/kafka"
	"github.com/wudian/wx/mongo"
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
	pair := goex.NewCurrencyPair2("ETH-BTC")
	goexTicker, err := global.Apis[config.API_OKEX].GetTicker(pair)
	jsonStr, err := utils.Struct2JsonString(goexTicker)
	if err == nil {
		t.Log(string(jsonStr))
	}

	ticker := utils.GoexTicker2Ticker(goexTicker, config.API_OKEX)
	jsonStr, err = utils.Struct2JsonString(ticker)
	if err == nil {
		t.Log(string(jsonStr))
	}

	client, err := mongo.NewMgoClient()
	if err==nil{
		client.Insert(ticker)
	}
}

func testKafka(t *testing.T)  {
	kafka.SyncProducer()
	assert.Equal(t, "123", "123")
}

