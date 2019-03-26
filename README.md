1.folder runner1 is previous , it's not used
global = &Global{
		apiNames: []string{"huobi", "okex", "hitbtc", "binance", "bithumb"}, // each exchange
		vecSymbols: []string{"btc_usdt", "eth_usdt"},// each symbol
	}

2.config file is global.go

3.we use timer to get ticker from ecah exchange, then calculate weighted mean value, and write it to kafka

4.server.go provide restful or websocket server

5.models.go provide self define struct, such as Ticker