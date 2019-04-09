package config

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

const (
	// "binance", "bithumb"  "huobi", , "hitbtc"
	API_HASHKEY = "hashkey"
	API_HUOBI = "huobi"
	API_OKEX = "okex"
	API_BINANCE = "binance"
	API_HITBTC = "hitbtc"
	API_BITHUMB = "bithumb"
)

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
	ApiNames []string
	VecSymbols []string
	// api -> symbol -> Ticker
	Tickers map[string]map[string]*goex.Ticker

	// api -> Weight
	Weight map[string]float64
	// weighted mean , symbol -> Ticker
	WeightMeanTickers map[string]*goex.Ticker
	MutexWeightMeanTickers sync.Mutex

	//api_name -> api
	Apis map[string]goex.API

	Duration uint64 // duration between api.date and local, out of range is invalid
	IsPrint bool // whether print log
	Log *log.Logger
	IsStoreData bool
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
		// "binance", "bithumb"   "huobi",
		ApiNames: []string{API_OKEX, API_HITBTC, API_BINANCE}, //exchange
		VecSymbols: []string{"BTC-USDT", "ETH-USDT", "ETH-BTC"},//
		Duration: 10,
		IsPrint: true,
		IsStoreData: true,
	}
	fileName := "wx.log"
	logFile,err  := os.Create(fileName)
	//defer logFile.Close()
	if err != nil {
		log.Fatalf("open file %s error !\n", fileName)
	}
	global.Log = log.New(logFile,"", log.Ltime)

	global.Tickers = map[string]map[string]*goex.Ticker{}
	for _, api := range global.ApiNames{
		global.Tickers[api] = map[string]*goex.Ticker{}
	}

	global.Weight = map[string]float64{}
	global.Apis = map[string]goex.API{}
	for _, api := range global.ApiNames{
		if api == API_HUOBI{
			global.Weight[api] = 1
			global.Apis[api] = huobi.NewHuoBiProSpot(httpProxyClient, apikey_huobi, secretkey_huobi)
		} else if api == API_OKEX{
			global.Weight[api] = 1
			global.Apis[api] = okcoin.NewOKExSpot(http.DefaultClient, apikey_okex, secretkey_okex)
		} else if api == API_HITBTC{
			global.Weight[api] = 1
			global.Apis[api] = hitbtc.New(http.DefaultClient, apikey_hitbtc, secretkey_hitbtc)
		} else if api == API_BINANCE{
			global.Weight[api] = 1
			global.Apis[api] = binance.New(http.DefaultClient, apikey_binance, secretkey_binance)
		} else if api == API_BITHUMB{
			global.Weight[api] = 0
			global.Apis[api] = bithumb.New(http.DefaultClient, apikey_bithumb, secretkey_bithumb)
		} else {
			panic("need add api")
		}
	}
	global.WeightMeanTickers = map[string]*goex.Ticker{}

	return global
}
