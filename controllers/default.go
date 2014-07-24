package controllers

import (
	"github.com/astaxie/beego"
	"ria/components"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	assets := components.Assets()

	if css, err := assets.Css(); err != nil {
		beego.Error(err)
	} else {
		this.Data["Css"] = css
	}

	if js, err := assets.Js(); err != nil {
		beego.Error(err)
	} else {
		this.Data["Js"] = js
	}

	this.Data["Storage"] = "{}"


	this.TplNames = "index.tpl"
}
