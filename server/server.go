package server

import (
	"github.com/astaxie/beego"
	"github.com/wudian/wx/config"
	"github.com/wudian/wx/utils"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	global := config.GlobalInstance()
	symbol := strings.ToUpper(this.GetString("symbol"))
	global.MutexWeightMeanTickers.Lock()
	defer global.MutexWeightMeanTickers.Unlock()
	jsonStr, err := utils.Struct2JsonString(utils.GoexTicker2Ticker(global.WeightMeanTickers[symbol]))
	if err == nil {
		this.Ctx.WriteString(jsonStr)
	}
}

func StartServer()  {
	// http://127.0.0.1:8080/market/ticker/?symbol=btc_usdt
	beego.Router("/market/ticker", &MainController{})
	beego.Run()
}
