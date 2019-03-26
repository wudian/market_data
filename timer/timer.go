package timer

import (
	"github.com/astaxie/beego/toolbox"
	"github.com/nntaoli-project/GoEx"
	"github.com/wudian/wx/config"
	"github.com/wudian/wx/utils"
	"math"
	"time"
)

func StartTimer() error {
	var (
		err error
		global = config.GlobalInstance()
	)
	tk := toolbox.NewTask("task", "0/5 * * * * *", func() error {

		now_second := time.Now().Unix()
		tmp_weight := global.Weight
		for _, api := range global.ApiNames {
			for _, symbol := range global.VecSymbols {
				pair := goex.NewCurrencyPair2(symbol)
				global.Tickers[api][symbol], err = global.Apis[api].GetTicker(pair)

				if err != nil {
					tmp_weight[api] = 0
					continue
				} else {
					dura := uint64(math.Abs(float64(global.Tickers[api][symbol].Date-now_second)))
					if dura>global.Duration{
						tmp_weight[api] = 0
						continue
					}
				}

				if global.IsPrint {
					jsonStr, err := utils.Struct2JsonString(utils.GoexTicker2Ticker(global.Tickers[api][symbol]))
					if err == nil {
						//t := time.Unix(global.Tickers[api][symbol].Date,0).Format("2006-01-02 15:04:05")
						global.Log.Printf("%s %s\n", api, jsonStr)
					}
				}
			}
		}

		global.MutexWeightMeanTickers.Lock()
		defer global.MutexWeightMeanTickers.Unlock()
		for _, symbol := range global.VecSymbols{
			sumTicker := goex.NewTicker()
			for _, api := range global.ApiNames {
				if global.Weight[api] >0 {
					sumTicker.Add(global.Tickers[api][symbol].Multi(global.Weight[api]))
				}

			}
			sumWei := float64(0)
			for _, wei := range global.Weight{
				sumWei += wei
			}
			sumTicker.Date = now_second
			sumTicker.Pair = goex.NewCurrencyPair2(symbol)
			global.WeightMeanTickers[symbol] = sumTicker.Div(sumWei).Decimal()

			if global.IsPrint {
				jsonStr, err := utils.Struct2JsonString(utils.GoexTicker2Ticker(global.WeightMeanTickers[symbol]))
				if err == nil {
					global.Log.Printf("weighted mean %s\n", jsonStr)
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