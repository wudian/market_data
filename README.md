1.version1 and version2 is previous version, they are not be used

2.config.go
subscribe each symbols from each exchange
global = &Global{
		ApiNames: []string{"huobi", "okex", "hitbtc", "binance", "bithumb"}, //exchange
		VecSymbols: []string{"BTC_USDT", "ETH_USDT"},//
	}

3.we use timer(beego.toolbox.task) to get ticker from ecah exchange, then calculate weighted mean value, and use timer to write it to websocket clients or kafka

4.server.go provide restful or websocket server base on beego

5.models.go provide self define struct, such as Ticker