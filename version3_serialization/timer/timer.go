package timer

import (
	"github.com/astaxie/beego/toolbox"
	"github.com/wudian/GoEx"
	"github.com/wudian/wx/config"
	"github.com/wudian/wx/server"
	"github.com/wudian/wx/utils"
	"log"
	"math"
	"sync"
	"time"
)

func StartTimer() error {
	var (
		err error
		global = config.GlobalInstance()
		// we need a lock to protect global.Tickers and global.WeightMeanTickers
		rdMutex sync.RWMutex
	)
	tk1 := toolbox.NewTask("task1", "0/1 * * * * *", func() error {
		rdMutex.Lock()
		defer rdMutex.Unlock()

		now_second := time.Now().Unix()
		tmp_weight := utils.DeepCopy(global.Weight).(map[string]float64)
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
						log.Printf("%s %s\n", api, jsonStr)
					}
				}
			}
		}

		global.MutexWeightMeanTickers.Lock()
		defer global.MutexWeightMeanTickers.Unlock()
		for _, symbol := range global.VecSymbols{
			sumTicker := goex.NewTicker()
			for _, api := range global.ApiNames {
				if tmp_weight[api] >0 {
					sumTicker.Add(global.Tickers[api][symbol].Multi(tmp_weight[api]))
				}

			}
			sumWei := float64(0)
			for _, wei := range tmp_weight{
				sumWei += wei
			}
			if sumWei == float64(0){
				break
			}
			sumTicker.Date = now_second
			sumTicker.Pair = goex.NewCurrencyPair2(symbol)
			global.WeightMeanTickers[symbol] = sumTicker.Div(sumWei).Decimal()

			if global.IsPrint {
				jsonStr, err := utils.Struct2JsonString(utils.GoexTicker2Ticker(global.WeightMeanTickers[symbol]))
				if err == nil {
					log.Printf("weighted mean %s\n", jsonStr)
				}
			}
		}

		server.SendTicker()
		return nil
	})
	err = tk1.Run()
	if err != nil {
		return err
	} else {
		toolbox.AddTask("task1", tk1)
	}

	toolbox.StartTask()
	//toolbox.StopTask()
	return nil
}