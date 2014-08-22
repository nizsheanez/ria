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

	user, err := models.FindUser(int(id))
	if err != nil {
		beego.Error(err)
	}

//	data, err := user.GetInitialData()
//	if err != nil {
//		beego.Error(err)
//	}
//
//	str, err := json.Marshal(data)
//	if err != nil {
//		beego.Error(err)
//	}
//	beego.Info(string(str))

	result, err := user.Get(int(id))
	if err != nil {
		this.Data["json"] = err
		this.ServeJson()
		return
	}
	this.Data["json"] = result

	this.ServeJson()
}
