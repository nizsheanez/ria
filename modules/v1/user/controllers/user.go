package controllers

import (
	"github.com/astaxie/beego"
	"ria/modules/v1/user/models"
	"errors"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Get() {
	var id int
	this.Ctx.Input.Bind(&id, "id")

	if id <= 0 {
		this.Data["json"] = errors.New("User id is required")
		this.ServeJson()
		return
	}

	user := models.NewUser()
	err := user.FindById(id)
	if err != nil {
		beego.Error(err)
	}

	this.Data["json"] = user

	this.ServeJson()
}


func (this *UserController) List() {
	users, err := models.FindUsers()

	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["json"] = users
	this.ServeJson()
}
