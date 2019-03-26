package wx

import (
	"encoding/json"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	global := GlobalInstance()
	symbol := this.GetString("symbol")
	global.mutexWeightMeanTickers.Lock()
	defer global.mutexWeightMeanTickers.Unlock()
	jsonStr, err := json.Marshal(global.weightMeanTickers[symbol])
	if err == nil {
		this.Ctx.WriteString(string(jsonStr))
	}
}

func StartServer()  {
	// http://127.0.0.1:8080/market/ticker/?symbol=btc_usdt
	beego.Router("/market/ticker", &MainController{})
	beego.Run()
}
