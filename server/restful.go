package server

import (
	"github.com/astaxie/beego"
	"github.com/wudian/market_data/config"
	"github.com/wudian/market_data/utils"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	global := config.GlobalInstance()
	symbol := strings.ToUpper(this.GetString("symbol"))
	global.RdMutex.RLock()
	defer global.RdMutex.RUnlock()
	jsonStr, err := utils.Struct2JsonString(utils.GoexTicker2Ticker(global.WeightMeanTickers[symbol], config.API_HASHKEY))
	if err == nil {
		this.Ctx.WriteString(jsonStr)
	}
}

func StartServer()  {
	beego.Router("/market/ticker", &MainController{})

	InitWebsocket()

	beego.Run()
}
