package controllers

import (
	"github.com/astaxie/beego"
	"ria/modules/v1/article/models"
	"errors"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	var id int
	this.Ctx.Input.Bind(&id, "id")

	if id <= 0 {
		this.Data["json"] = errors.New("User id is required")
		this.ServeJson()
		return
	}

	article, err := models.NewArticle().FindById(id)
	if err != nil {
		beego.Error(err)
	}

	this.Data["json"] = article

	this.ServeJson()
}

func (this *ArticleController) Post() {
	article := models.NewArticle()
	if err := this.ParseForm(article); err != nil {
		//some error
	} else {
		ok, err := article.Validate("create");
		if ok {
			this.Data["json"] = "ok"
		} else {
			this.Data["json"] = err
		}

	}
	this.ServeJson()
}

func (this *ArticleController) List() {
	this.ServeJson()
}
