package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"ria/conf"
	"encoding/json"
	"ria/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Css"] = conf.GetCss()
	this.Data["Js"] = conf.GetJs()

	user, err := models.FindUser(int(1))
	if err != nil {
		beego.Error(err)
	}

	data, err := user.GetInitialData()
	if err != nil {
		beego.Error(err)
	}

	str, err := json.Marshal(map[string]interface {}{"init":data})
	if err != nil {
		beego.Error(err)
	}

	this.Data["Storage"] = template.JS(string(str))

	this.TplNames = "index.tpl"
	return
}
