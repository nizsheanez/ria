package models

import (
	"ria/models"
	"ria/components"
	"github.com/lann/squirrel"
	"github.com/astaxie/beego/validation"
)

type Article struct {
	components.BaseModel
	Id                       int                    `db:"id" json:"id"`
	Title                    models.NullString      `db:"title" json:"title"`
	Description              models.NullString      `db:"description" json:"description"`
	CreateTime               string                 `db:"create_time" json:"create_time"`
	UpdateTime               string                 `db:"update_time" json:"update_time"`
}

func init() {

}

func (this *Article) FindById(id int) (*Article, error) {
	qb := squirrel.Select("*").From("model").Where("id = ?", id)

	err := components.LoadOne(&qb, &this)
	if err != nil {
		return nil, err
	}

	return this, nil
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

