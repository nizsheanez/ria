package models

import "github.com/astaxie/beego/orm"

type GoalCategory struct {
	Id                 int              `orm:"pk;auto;column(id)" json:"id"`
	Name               string           `json:"name"`
	CreateTime         string           `json:"create_time"`
	UpdateTime         string           `json:"update_time"`
	Goals              []*Goal          `orm:"reverse(many)" json:"-"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(GoalCategory))
}

