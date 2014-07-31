package main

import (
	_ "ria/routers"
	"github.com/astaxie/beego"
	"ria/controllers"
	"net/http"
)

func main() {
	//websocket server
	wampServer := controllers.NewServer()
	go wampServer.ListenAndServe()

	//run upgrade server, for upgrade http to ws
	//TODO: maybe it better implement as controller
	addr := ":8081"
	upgradeServer := &http.Server{
		Addr: addr,
		Handler: wampServer,
	}

	beego.BeeLogger.Info("Running Upgrade server on %s", addr)
	go upgradeServer.ListenAndServe()

	//run usual web app
	beego.Run()
}

