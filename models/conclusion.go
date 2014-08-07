package models

import "github.com/astaxie/beego/orm"

type Conclusion struct {
	Id                      int            `orm:"pk;auto;column(id)" json:"id"`
	Description             string         `json:"description"`
	ReportDate              string         `json:"report_date"`
	CreateTime              string         `json:"create_time"`
	UpdateTime              string         `json:"update_time"`
	User                 	*User          `orm:"rel(one);column(fk_user)"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Conclusion))
}
