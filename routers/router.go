package routers

import (
	"ria/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user/get", &controllers.UserController{})
}
