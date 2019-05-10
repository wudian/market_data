package main

import (
	"github.com/okcoin-okex/open-api-v3-sdk/okex-go-sdk-api"
	"strings"
)

const (
	HUOBI = 0
	OKEX  = 1

	APINUM = 5
)

type Global struct {
	apiNames    []string
	vecSymbols  []string
	tickers     map[string]map[string]interface{}
	client_okex *okex.Client

	isPrint bool
}

func GlobalInstance() *Global {
	var global Global
	global.apiNames = []string{"huobi", "okex"}
	global.vecSymbols = []string{"eth_usdt", "eth_btc"}
	global.tickers = map[string]map[string]interface{}{}
	for _, api := range global.apiNames {
		global.tickers[api] = map[string]interface{}{}
	}

	global.client_okex = okex.NewTestClient()
	global.isPrint = true
	return &global
}

func SymbolAdaptToApi(api, symbol string) string {
	symbol_api := symbol
	if api == "huobi" {
		symbol_api = strings.ReplaceAll(symbol, "_", "")
	}

	return symbol_api
}

func ApiNameAdaptToEnum(api string) int {
	var no int
	if api == "huobi" {
		no = HUOBI
	} else if api == "okex" {
		no = OKEX
	}
	return no
}
