package main

import (
	"github.com/astaxie/beego"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	global := GlobalInstance()
	symbol := strings.ToUpper(this.GetString("symbol"))
	global.mutexWeightMeanTickers.Lock()
	defer global.mutexWeightMeanTickers.Unlock()
	jsonStr, err := Struct2JsonString(GoexTicker2Ticker(global.weightMeanTickers[symbol]))
	if err == nil {
		this.Ctx.WriteString(jsonStr)
	}
}

func StartServer()  {
	// http://127.0.0.1:8080/market/ticker/?symbol=btc_usdt
	beego.Router("/market/ticker", &MainController{})
	beego.Run()
}
