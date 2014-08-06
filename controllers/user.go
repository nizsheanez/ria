package controllers

import (
	"github.com/astaxie/beego"
	"ria/models"
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

	user := &models.User{}
	result, err := user.Get(int(id))
	if err != nil {
		this.Data["json"] = err
		this.ServeJson()
		return
	}
	this.Data["json"] = result
	this.ServeJson()
}
