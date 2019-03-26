package server

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"github.com/wudian/wx/config"
	"github.com/wudian/wx/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	clients   = make(map[*websocket.Conn]bool)
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


func (this *MyWebSocketController) Get() {

	ws, err := upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		log.Println(err)
		this.Ctx.WriteString(err.Error())
		return
	}
	//  defer ws.Close()

	clients[ws] = true

	global := config.GlobalInstance()
	symbol := strings.ToUpper(this.GetString("symbol"))
	//不断的广播发送到页面上
	for {
		//目前存在问题 定时效果不好 需要在业务代码替换时改为beego toolbox中的定时器
		time.Sleep(time.Second * 3)
		global.MutexWeightMeanTickers.Lock()
		defer global.MutexWeightMeanTickers.Unlock()
		jsonStr, err := utils.Struct2JsonString(utils.GoexTicker2Ticker(global.WeightMeanTickers[symbol]))
		if err == nil {
			broadcast <- jsonStr
		}

	}
}

func InitWebsocket() {
	beego.Router("/ws", &MyWebSocketController{})

	go handleMessages()
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