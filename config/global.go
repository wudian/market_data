package config

import (
	"github.com/nntaoli-project/GoEx"
	"github.com/nntaoli-project/GoEx/binance"
	"github.com/nntaoli-project/GoEx/bithumb"
	"github.com/nntaoli-project/GoEx/hitbtc"
	"github.com/nntaoli-project/GoEx/huobi"
	"github.com/nntaoli-project/GoEx/okcoin"
	"github.com/wonderivan/logger"
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
	// name -> url
	ApiNames map[string]string
	VecSymbols []string
	// api -> symbol -> Ticker
	Tickers map[string]map[string]*goex.Ticker

	// api -> Weight
	Weight map[string]float64
	// weighted mean , symbol -> Ticker
	WeightMeanTickers map[string]*goex.Ticker

	// we need a lock to protect global.Tickers and global.WeightMeanTickers
	RdMutex sync.RWMutex

	//api_name -> api
	Apis map[string]goex.API

	Duration uint64 // duration between api.date and local, out of range is invalid
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
	v, err := readXml()
	if err != nil {
		logger.Alert(err)
		os.Exit(1)
	}
	global = &Global{
		// "binance", "bithumb"   "huobi",
		ApiNames: map[string]string{}, //exchange
		VecSymbols: []string{},//"BTC-USDT", "ETH-USDT",
		Duration: 10,
		IsStoreData: false,
	}
	setGlobal(v)

	global.Tickers = map[string]map[string]*goex.Ticker{}
	for api, _ := range global.ApiNames{
		global.Tickers[api] = map[string]*goex.Ticker{}
	}

	global.Weight = map[string]float64{}
	global.Apis = map[string]goex.API{}
	for name, url := range global.ApiNames{
		if name == API_HUOBI{
			global.Weight[name] = 1
			huobi.API_BASE_URL = url
			global.Apis[name] = huobi.NewHuoBiProSpot(http.DefaultClient, apikey_huobi, secretkey_huobi)
		} else if name == API_OKEX{
			global.Weight[name] = 1
			okcoin.API_BASE_URL = url
			global.Apis[name] = okcoin.NewOKExSpot(http.DefaultClient, apikey_okex, secretkey_okex)
		} else if name == API_HITBTC{
			global.Weight[name] = 1
			hitbtc.API_BASE_URL = url
			global.Apis[name] = hitbtc.New(http.DefaultClient, apikey_hitbtc, secretkey_hitbtc)
		} else if name == API_BINANCE{
			global.Weight[name] = 1
			binance.API_BASE_URL = url
			global.Apis[name] = binance.New(http.DefaultClient, apikey_binance, secretkey_binance)
		} else if name == API_BITHUMB{
			global.Weight[name] = 0
			global.Apis[name] = bithumb.New(http.DefaultClient, apikey_bithumb, secretkey_bithumb)
		} else {
			panic("need add api")
		}
	}
	global.WeightMeanTickers = map[string]*goex.Ticker{}

	return global
}

