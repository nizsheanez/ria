package models

import "github.com/astaxie/beego/orm"

type Report struct {
	Id                      int            `orm:"pk;auto;column(id)" json:"id"`
	Description             string         `json:"description"`
	ReportDate              string         `json:"report_date"`
	CreateTime              string         `json:"create_time"`
	UpdateTime              string         `json:"update_time"`
	Goal                 	*Goal 		   `orm:"rel(one);column(fk_goal)"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Report))
}
