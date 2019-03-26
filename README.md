# wx
get market data from other exchange and calculate the weighted mean value, then publish it

1.folder runner1 is previous , it's not used
global = &Global{
		apiNames: []string{"huobi", "okex", "hitbtc", "binance", "bithumb"}, // each exchange
		vecSymbols: []string{"btc_usdt", "eth_usdt"},// each symbol
	}

2.config file is global.go

3.we use timer to get ticker from ecah exchange, then calculate weighted mean value, and write it to kafka

4.Server.go privde restful server, but now not used
