package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//symbol := this.GetString("symbol")
	//jsonStr, err := json.Marshal(tickers[symbol])
	//if err == nil {
	//	this.Ctx.WriteString(string(jsonStr))
	//}
}

func StartServer()  {
	//beego.Router("/market/detail/merged", &MainController{})
	beego.Run()
}
