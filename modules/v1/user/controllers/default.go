package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"ria/conf"
	"encoding/json"
	"ria/modules/v1/user/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Css"] = conf.GetCss()
	this.Data["Js"] = conf.GetJs()

	user := models.NewUser()
	err := user.FindById(int(1))
	if err != nil {
		beego.Error(err)
		return
	}

	data, err := user.GetInitialData()
	if err != nil {
		beego.Error(err)
	}

	str, err := json.Marshal(map[string]interface{}{"init":data})
	if err != nil {
		beego.Error(err)
	}

	this.Data["Storage"] = template.JS(string(str))

	this.TplNames = "/index.tpl"
	return
}
