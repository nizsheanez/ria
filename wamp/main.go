package main

import (
	"github.com/astaxie/beego"
	"net/http"
	"ria/wamp/components"
)

func main() {
	//websocket server
	wampServer := components.NewServer()
	go wampServer.ListenAndServe()

	//run upgrade server, for upgrade http to ws
	addr := ":8081"
	upgradeServer := &http.Server{
		Addr: addr,
		Handler: wampServer,
	}


	beego.BeeLogger.Info("Running Upgrade server on %s", addr)
	upgradeServer.ListenAndServe()
}

