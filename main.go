package main

import (
	_ "ria/routers"
	"github.com/astaxie/beego"
	"ria/controllers"
)

func main() {
	// websocket server
	server := controllers.NewServer(":8081")
	go server.Listen()

	beego.Run()
}

