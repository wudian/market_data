package main

import (
	"github.com/nntaoli-project/GoEx"
	"github.com/nntaoli-project/GoEx/binance"
	"github.com/nntaoli-project/GoEx/bithumb"
	"github.com/nntaoli-project/GoEx/hitbtc"
	"github.com/nntaoli-project/GoEx/huobi"
	"github.com/nntaoli-project/GoEx/okcoin"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

//const (
//	HUOBI = 0
//	OKEX = 1
//
//	APINUM = 5
//)

var httpProxyClient = &http.Client{
	Transport: &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return &url.URL{
				Scheme: "socks5",
				Host:   "127.0.0.1:55307"}, nil
		},
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
	},
	Timeout: 10 * time.Second,
}

var (
	apikey_huobi    = "478c732a-4f0cd03e-e779076d-0a87b"
	secretkey_huobi = "e1b9b5c7-5a9bff84-e72dc5eb-f5dd9"

	apikey_okex    = "06aa16ef-ec6e-425b-bbdb-145bee989842"
	secretkey_okex = "EF674EDE16BE81D7F2A2F83330BEDE36"

	apikey_binance    = "zRuNMVWQrFRswwaccIlZ9liNrMlJboEKnopZMfyjaEvAvHdbI8F8tg4wH7ruRpEv"
	secretkey_binance = "vTSIXM9xNyOMDMgvrML4GPGKfFpnY45f89xUoBayw2xC8b15NND1HSp8e9mXX5py"

	apikey_bithumb    = ""
	secretkey_bithumb = ""

	apikey_hitbtc    = ""
	secretkey_hitbtc = ""
)

type Global struct {
	apiNames []string
	vecSymbols []string
	// api -> symbol -> Ticker
	tickers map[string]map[string]*goex.Ticker

	// api -> weight
	weight map[string]float64
	// weighted mean , symbol -> Ticker
	weightMeanTickers map[string]*goex.Ticker
	mutexWeightMeanTickers sync.Mutex

	//api_name -> api
	apis map[string]goex.API

	duration uint64 // duration between api.date and local, out of range is invalid
	isPrint bool // whether print log
	log *log.Logger
}

var global *Global
var mu sync.Mutex
func GlobalInstance() *Global {
	mu.Lock()
	defer mu.Unlock()
	if global != nil{
		return global
	}
	global = &Global{
		apiNames: []string{"huobi", "okex", "hitbtc", "binance", "bithumb"}, //,
		vecSymbols: []string{"BTC_USDT", "ETH_USDT"},//
		duration: 10,
		isPrint: true,
	}
	fileName := "wx.log"
	logFile,err  := os.Create(fileName)
	//defer logFile.Close()
	if err != nil {
		log.Fatalf("open file %s error !\n", fileName)
	}
	global.log = log.New(logFile,"", log.Ltime)

	global.tickers = map[string]map[string]*goex.Ticker{}
	for _, api := range global.apiNames{
		global.tickers[api] = map[string]*goex.Ticker{}
	}

	global.weight = map[string]float64{}
	global.apis = map[string]goex.API{}
	for _, api := range global.apiNames{
		if api == "huobi"{
			global.weight[api] = 1
			global.apis[api] = huobi.NewHuoBiProSpot(httpProxyClient, apikey_huobi, secretkey_huobi)
		} else if api == "okex"{
			global.weight[api] = 1
			global.apis[api] = okcoin.NewOKExSpot(http.DefaultClient, apikey_okex, secretkey_okex)
		} else if api == "hitbtc"{
			global.weight[api] = 1
			global.apis[api] = hitbtc.New(http.DefaultClient, apikey_hitbtc, secretkey_hitbtc)
		} else if api == "binance"{
			global.weight[api] = 1
			global.apis[api] = binance.New(http.DefaultClient, apikey_binance, secretkey_binance)
		} else if api == "bithumb"{
			global.weight[api] = 0
			global.apis[api] = bithumb.New(http.DefaultClient, apikey_bithumb, secretkey_bithumb)
		} else {
			panic("need add api")
		}
	}
	global.weightMeanTickers = map[string]*goex.Ticker{}

	return global
}
