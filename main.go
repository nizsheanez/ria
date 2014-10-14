package main

import (
	_ "ria/routers"
	"github.com/astaxie/beego"
	"net/http"
	components "ria/components"
	wamp "ria/wamp/components"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	beego.BeeLogger.Info("!!!!!!")

	db, err := sqlx.Open("mysql", "root:asharov@tcp(db:3306)/blog3?charset=utf8")
	if err != nil {
		panic(err)
	}
	components.App.Db = db

	//websocket server
	wampServer := wamp.NewServer()
	beego.BeeLogger.Info("Running WS server wait for client requests")

	go wampServer.ListenAndServe()

	//run upgrade server, for upgrade http to ws
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
