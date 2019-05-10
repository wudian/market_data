package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"github.com/huobiapi/REST-GO-demo/services"
)

func StartTimer(global Global) error {
	var err error
	tk := toolbox.NewTask("task", "0/5 * * * * *", func() error {
		for _, api := range global.apiNames {
			for _, symbol := range global.vecSymbols {
				symbol_api := SymbolAdaptToApi(api, symbol)
				if api == "huobi" {
					global.tickers[api][symbol] = services.GetTicker(symbol_api)
				} else if api == "okex" {
					global.tickers[api][symbol], err = global.client_okex.GetSpotInstrumentTicker(symbol_api)
				}

				if global.isPrint {
					jsonStr, err := json.Marshal(global.tickers[api][symbol])
					if err == nil {
						fmt.Println(api + string(jsonStr))
					}
				}
			}
		}
		return nil
	})
	err = tk.Run()
	if err != nil {
		return err
	} else {
		toolbox.AddTask("task", tk)
	}
	toolbox.StartTask()
	//toolbox.StopTask()
	return nil
}
