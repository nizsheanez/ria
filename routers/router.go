package routers

import (
	"ria/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &controllers.UserController{})
    beego.Router("/user/list", &controllers.UserController{}, "get:List")
}
