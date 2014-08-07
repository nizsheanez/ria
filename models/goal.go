package models

import "github.com/astaxie/beego/orm"

type Goal struct {
	Id                 int                      `orm:"pk;auto;column(id)" json:"id"`
	Title              string                  `json:"title"`
	Status             int                      `json:"status"`
	Completed          int                      `json:"password_hash"`
	Reason             string              		`json:"reason"`
	CreateTime         string                  `json:"create_time"`
	UpdateTime         string                  `json:"update_time"`
	Decomposition      string                  `json:"decomposition"`
	Comments           string              		`json:"comments"`
	Category           *GoalCategory          `orm:"rel(fk);column(fk_goal_category)"`
	User               *User                  `orm:"rel(fk);column(fk_user)"`
	Today              *Day          			`orm:"-" json:"today"`
	Yesterday          *Day          			`orm:"-" json:"yesterday"`
	//	Steps              []*Step                `orm:"reverse(many)"`
}

type Day struct {
	Reports            []*Report     `orm:"-" json:"report"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Goal))
}

