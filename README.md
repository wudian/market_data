1.version1 and version2 is previous version, they are not be used

2.config.go
subscribe each symbols from each exchange
global = &Global{
		ApiNames: []string{"huobi", "okex", "hitbtc", "binance", "bithumb"}, //exchange
		VecSymbols: []string{"BTC_USDT", "ETH_USDT"},//
	}

3.we use timer(beego.toolbox.task) to get ticker from ecah exchange, then calculate weighted mean value, and use timer to write it to websocket clients or kafka

4.server.go provide restful or websocket server base on beego
restful:   http://127.0.0.1:8080/market/ticker/?symbol=btc_usdt
websocket:   https://www.bejson.com/httputil/websocket/     ws://127.0.0.1:8080/ws    {"symbol":"btc_usdt"}
kafka:

5.models.go provide self define struct, such as Ticker

