package routers

import (
	"ria/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

	// WebSocket.
	beego.Router("/ws", &controllers.WebSocketController{})
	beego.Router("/ws/join", &controllers.WebSocketController{}, "get:Join")
	beego.Router("/ws/user", &controllers.WebSocketController{}, "get:User")
}
