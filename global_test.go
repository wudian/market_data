package wx

import (
	"encoding/json"
	"github.com/nntaoli-project/GoEx"
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
	ret, err := global.apis["binance"].GetTicker(goex.BTC_USDT)
	jsonStr, err := json.Marshal(ret)
	if err == nil {
		t.Log(string(jsonStr))
	}
}

func testKafka(t *testing.T)  {
	syncProducer()
	assert.Equal(t, "123", "123")
}