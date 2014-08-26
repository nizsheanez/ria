package controllers

import (
	"github.com/astaxie/beego"
	"ria/models"
	"ria/conf"
)

type NewsController struct {
	beego.Controller
}

func (this *NewsController) Get() {
	this.Data["Css"] = conf.GetCss()
	this.Data["Js"] = conf.GetJs()

	users, err := models.FindUsers()

	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["json"] = users
	this.ServeJson()
}
