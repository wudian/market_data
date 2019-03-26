package server

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/wudian/wx/config"
	"github.com/wudian/wx/models"
	"github.com/wudian/wx/utils"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	clients   = make(map[*websocket.Conn][]string)
	mutexClients sync.Mutex

	ticker = time.NewTicker(time.Second * 3)

	broadcast = make(chan string)
	upgrader = websocket.Upgrader{
		// 解决跨域问题
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type MyWebSocketController struct {
	beego.Controller
}

func disconnect(ws *websocket.Conn) {
	mutexClients.Lock()
	defer mutexClients.Unlock()
	delete(clients, ws)
	ws.Close()
}

func (this *MyWebSocketController) Get() {

	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		//log.Println(err)
		this.Ctx.WriteString(err.Error())
		return
	}
	defer disconnect(ws)

	//global := config.GlobalInstance()

	clients[ws] = []string{}
	times := 0
	//不断的广播发送到页面上
	for {
		var subReq models.SubReq
		err = ws.ReadJSON(&subReq)
		if err != nil {
			times += 1
			this.Ctx.WriteString(err.Error())
			if times > 10{
				return
			}
			continue
		}
		symbol := strings.ToUpper(subReq.Symbol)
		mutexClients.Lock()
		defer mutexClients.Unlock()
		clients[ws] = append(clients[ws], symbol)


		//目前存在问题 定时效果不好 需要在业务代码替换时改为beego toolbox中的定时器
		//time.Sleep(time.Second * 3)
		//global.MutexWeightMeanTickers.Lock()
		//defer global.MutexWeightMeanTickers.Unlock()
		//jsonStr, err := utils.Struct2JsonString(utils.GoexTicker2Ticker(global.WeightMeanTickers[symbol]))
		//if err == nil {
		//	broadcast <- jsonStr
		//}

	}
}

func InitWebsocket() {
	//ws://127.0.0.1:8080/ws
	beego.Router("/ws", &MyWebSocketController{})

	//go handleMessages()
	go sendTicker()
}

func sendTicker() {
	global := config.GlobalInstance()
	for {
		select {
		case <-ticker.C:
			//fmt.Printf("ticked at %v\n", time.Now())
			//mutexClients.Lock()
			//defer mutexClients.Unlock()
			for client, vecSymbols := range clients{
				for _, symbol := range vecSymbols{
					jsonStr, err := utils.Struct2JsonString(utils.GoexTicker2Ticker(global.WeightMeanTickers[symbol]))
					if err == nil {
						err := client.WriteJSON(jsonStr)
						if err != nil {
							//log.Printf("client.WriteJSON error: %v", err)
							disconnect(client)
						}
					}
				}
			}
		}
	}
}

//广播发送至页面
func handleMessages() {
	for {
		msg := <-broadcast
		fmt.Println("clients len ", len(clients))
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}