package components

import (
	"github.com/astaxie/beego"
	"github.com/jmoiron/sqlx"
)

type Application struct {
	beego.App
	Db *sqlx.DB
}

var (
	App *Application
)

func init() {
	App = &Application{}
}

