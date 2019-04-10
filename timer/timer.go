package timer

import (
	"github.com/astaxie/beego/toolbox"
	"github.com/nntaoli-project/GoEx"
	"github.com/wudian/wx/config"
	"github.com/wudian/wx/mongo"
	"github.com/wudian/wx/server"
	"github.com/wudian/wx/utils"
	"log"
	"math"
	"time"
)

var (
	err    error
	global = config.GlobalInstance()
	// api -> weight
	tmpWeight = map[string]float64{}
	nowSecond int64
	mgoClient mongo.MgoClient
)

func GetTicker(api, symbol string) {
	for {
		//time.Sleep(1*time.Second)
		nowSecond = time.Now().Unix()
		pair := goex.NewCurrencyPair2(symbol)
		ticker, err := global.Apis[api].GetTicker(pair)
		global.RdMutex.Lock()
		if err != nil {
			tmpWeight[api] = 0
			global.Log.Println(err.Error())
		} else {
			dura := uint64(math.Abs(float64(ticker.Date - nowSecond)))
			if dura > global.Duration {
				tmpWeight[api] = 0
				t := time.Unix(nowSecond,0).Format("2006-01-02 15:04:05")
				global.Log.Printf("dura: %d, time:%s\n", dura, t)
			} else {
				tmpWeight[api] = global.Weight[api]
				global.Tickers[api][symbol] = ticker
			}
		}
		global.RdMutex.Unlock()
		myTicker := utils.GoexTicker2Ticker(ticker, api)
		if global.IsStoreData && err == nil {
			mgoClient.Insert(myTicker)
		}
		if global.IsPrint && err == nil {
			jsonStr, err := utils.Struct2JsonString(myTicker)
			if err == nil {
				log.Println(jsonStr)
			}
		}
	}

}

func StartTimer() error {
	mgoClient, err = mongo.NewMgoClient()
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	for _, api := range global.ApiNames {
		for _, symbol := range global.VecSymbols {
			go GetTicker(api, symbol)
		}
	}

	tk1 := toolbox.NewTask("task1", "0/1 * * * * *", func() error {
		global.RdMutex.Lock()
		defer global.RdMutex.Unlock()

		for _, symbol := range global.VecSymbols{
			sumTicker := goex.NewTicker()
			for _, api := range global.ApiNames {
				if tmpWeight[api] >0 {
					sumTicker.Add(global.Tickers[api][symbol].Multi(tmpWeight[api]))
				}

			}
			sumWei := float64(0)
			for _, wei := range tmpWeight{
				sumWei += wei
			}
			if sumWei == float64(0){
				break
			}
			sumTicker.Date = nowSecond
			sumTicker.Pair = goex.NewCurrencyPair2(symbol)
			global.WeightMeanTickers[symbol] = sumTicker.Div(sumWei).Decimal()

			if global.IsPrint {
				jsonStr, err := utils.Struct2JsonString(utils.GoexTicker2Ticker(global.WeightMeanTickers[symbol], config.API_HASHKEY))
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
