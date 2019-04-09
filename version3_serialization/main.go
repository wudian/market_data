package main

import (
	"github.com/wudian/wx/server"
	"github.com/wudian/wx/timer"
)

func main() {
	timer.StartTimer()
	server.StartServer()
}
