package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

type Goal struct {
	Id                 int              `orm:"pk;auto;column(id)" json:"id"`
	Title              string           `json:"title"`
	Status             int              `json:"status"`
	Completed          int              `json:"password_hash"`
	Reason             string           `json:"reason"`
	CreateTime         string           `json:"create_time"`
	UpdateTime         string           `json:"update_time"`
	Decomposition      string           `json:"decomposition"`
	Comments           string           `json:"comments"`
	FkGoalCategory     int              `json:"fk_goal_category"`
	User               *User            `orm:"rel(fk);column(fk_user)"`
}

type User struct {
	Id                   int            `orm:"pk" db:"id" json:"id"`
	UserName             string         `orm:"column(username)" db:"username" json:"username"`
	AuthKey              string         `json:"auth_key"`
	PasswordHash         string         `json:"password_hash"`
	PasswordResetToken   string         `json:"password_reset_token"`
	Email                string         `json:"email"`
	Role                 int            `json:"role"`
	Status               int8           `json:"status"`
	CreateTime           string         `json:"create_time"`
	UpdateTime           string         `json:"update_time"`
	Goals				[]*Goal 		`orm:"reverse(many)"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User), new(Goal))
}

func (this *User) View(arguments []interface{}) (map[string]interface{}, error) {
	o := orm.NewOrm()
	o.Using("default")

	user := &User{Id: 1}
	err := o.Read(user)

	o.LoadRelated(user, "Goals")

	if err != nil {
		return nil, err
	}

	data := map[string]interface{}{
//		"email":user.Email,
	}
	return data, nil
}

