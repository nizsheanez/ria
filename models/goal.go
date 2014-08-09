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
	Category           *GoalCategory          `orm:"rel(fk);column(fk_goal_category)" json:"-"`
	User               *User                  `orm:"rel(fk);column(fk_user)" json:"-"`
	//	Steps              []*Step                `orm:"reverse(many)"`
}

type Day struct {
	Report            *Report     `orm:"-" json:"report"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Goal))
}

