package routers

import (
	article "ria/modules/v1/article/controllers"
	user "ria/modules/v1/user/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &user.MainController{})
	beego.Router("/user", &user.UserController{})
	beego.Router("/user/list", &user.UserController{}, "get:List")
	beego.Router("/article/", &article.ArticleController{})
	beego.Router("/article/list", &article.ArticleController{}, "get:List")
}
