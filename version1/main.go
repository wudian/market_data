package main

var ()

func main() {
	//test3()
	//return

	global := GlobalInstance()
	StartTimer(*global)
	StartServer()
}
