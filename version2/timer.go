package main

import (
	"github.com/astaxie/beego/toolbox"
	"github.com/wudian/GoEx"
	"math"
	"time"
)

func StartTimer() error {
	var (
		err    error
		global = GlobalInstance()
	)
	tk := toolbox.NewTask("task", "0/5 * * * * *", func() error {

		now_second := time.Now().Unix()
		tmp_weight := global.weight
		for _, api := range global.apiNames {
			for _, symbol := range global.vecSymbols {
				pair := goex.NewCurrencyPair2(symbol)
				global.tickers[api][symbol], err = global.apis[api].GetTicker(pair)

				if err != nil {
					tmp_weight[api] = 0
					continue
				} else {
					dura := uint64(math.Abs(float64(global.tickers[api][symbol].Date - now_second)))
					if dura > global.duration {
						tmp_weight[api] = 0
						continue
					}
				}

				if global.isPrint {
					jsonStr, err := Struct2JsonString(GoexTicker2Ticker(global.tickers[api][symbol]))
					if err == nil {
						//t := time.Unix(global.tickers[api][symbol].Date,0).Format("2006-01-02 15:04:05")
						global.log.Printf("%s %s\n", api, jsonStr)
					}
				}
			}
		}

		global.mutexWeightMeanTickers.Lock()
		defer global.mutexWeightMeanTickers.Unlock()
		for _, symbol := range global.vecSymbols {
			sumTicker := goex.NewTicker()
			for _, api := range global.apiNames {
				if global.weight[api] > 0 {
					sumTicker.Add(global.tickers[api][symbol].Multi(global.weight[api]))
				}

			}
			sumWei := float64(0)
			for _, wei := range global.weight {
				sumWei += wei
			}
			sumTicker.Date = now_second
			sumTicker.Pair = goex.NewCurrencyPair2(symbol)
			global.weightMeanTickers[symbol] = sumTicker.Div(sumWei).Decimal()

			if global.isPrint {
				jsonStr, err := Struct2JsonString(GoexTicker2Ticker(global.weightMeanTickers[symbol]))
				if err == nil {
					global.log.Printf("weighted mean %s\n", jsonStr)
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
