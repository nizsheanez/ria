package routers

import (
	"ria/controllers"
	article "ria/modules/v1/article/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/user/list", &controllers.UserController{}, "get:List")
	beego.Router("/article/", &article.ArticleController{})
	beego.Router("/article/list", &article.ArticleController{}, "get:List")
}
