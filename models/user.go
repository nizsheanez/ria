package models

import "github.com/astaxie/beego/orm"

type User struct {
	Id                   int            `orm:"pk;auto;column(id)" json:"id"`
	UserName             string         `orm:"column(username)" db:"username" json:"username"`
	AuthKey              string         `json:"auth_key"`
	PasswordHash         string         `json:"password_hash"`
	PasswordResetToken   string         `json:"password_reset_token"`
	Email                string         `json:"email"`
	Role                 int            `json:"role"`
	Status               int8           `json:"status"`
	CreateTime           string         `json:"create_time"`
	UpdateTime           string         `json:"update_time"`
	Goals                []*Goal        `orm:"reverse(many)"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}

func (this *User) Get(id int) (result map[string]interface{}, err error) {
	o := orm.NewOrm()
	o.Using("default")

	user := &User{Id: id}
	err = o.Read(user)

	o.LoadRelated(user, "Goals")

	if err != nil {
		return nil, err
	}

	result = map[string]interface{}{
		"email":user.Email,
	}
	return result, nil
}

