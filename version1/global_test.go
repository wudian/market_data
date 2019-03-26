package main

import (
	"fmt"
	"testing"
)

func TestSymbolAdaptToApi(t *testing.T) {
	global := GlobalInstance()
	for _, api := range global.apiNames {
		for _, symbol := range global.vecSymbols {
			symbol_api := SymbolAdaptToApi(api, symbol)
			fmt.Println(symbol + "->"+api+"->"+symbol_api)
		}
	}

}