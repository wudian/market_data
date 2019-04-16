## 用go接huobi、okex、binance等交易所的ticker行情，并计算加权平均值，通过kafka/restful/websocket提供服务

0.set http_proxy=http://127.0.0.1:55304  https_proxy=http://127.0.0.1:55304

1.version1 and version2 is previous version, they are not be used

2.config.go
subscribe each symbols from each exchange
global = &Global{
		ApiNames: []string{"huobi", "okex", "hitbtc", "binance", "bithumb"}, //exchange
		VecSymbols: []string{"BTC_USDT", "ETH_USDT"},//
	}

3.we use timer(beego.toolbox.task) to get ticker from ecah exchange, then calculate weighted mean value, and use timer to write it to websocket clients or kafka

4.server.go provide restful or websocket server base on beego
restful:   http://127.0.0.1:8080/market/ticker/?symbol=btc-usdt
websocket:   https://www.bejson.com/httputil/websocket/     ws://127.0.0.1:8080/ws    {"symbol":"eth-btc"}

5.kafka

6.models.go provide self define struct, such as Ticker


#mongo win



#kafka win
1.install java sdk and add path

2.install zookeeper to C:\Program Files (x86)\Green\zookeeper-3.4.9
modify conf/zoo.cfg:dataDir=
add path
cmd: zkServer

3.install kafka	to C:\kafka_2.12-2.1.1
	attention: install dir not in C:\Program Files (x86) becasue of blank
modify config/server.properties:log.dirs=
cmd: .\bin\windows\kafka-server-start.bat .\config\server.properties

4.kafka-console-producer.bat --broker-list localhost:9092 --topic test0811

5.kafka-console-consumer.bat --bootstrap-server localhost:9092 --topic test0811 --from-beginning
