package models

import (
	"ria/components"
	"github.com/astaxie/beego/validation"
)

type Article struct {
	components.BaseModel
	Id                       int                    `db:"id" json:"id" form:"id"`
	Title                    string                    `db:"title" json:"title" form:"title"`
	Description              string                    `db:"description" json:"description" form:"description"`
	CreateTime               string                 `db:"create_time" json:"create_time" form:"create_time"`
	UpdateTime               string                 `db:"update_time" json:"update_time" form:"update_time"`
}

func init() {

}

func (*Article) TableName() string {
	return "article"
}

func NewArticle() (*Article) {
	return &Article{}
}

func (this *Article) Validators(validator *validation.Validation, scenario string) {
	if scenario == "create" || scenario == "update" {
		validator.Required(this.Description, "description");
		validator.Min(this.Description, 0, "description");
		validator.Max(this.Description, 12000, "description");
		validator.Max(this.Title, 3, "title");
		validator.Min(this.Title, 255, "title");
	}
}

