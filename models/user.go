package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id                   int                       `orm:"pk;auto;column(id)" json:"id"`
	UserName             string                    `orm:"column(username)" db:"username" json:"username"`
	AuthKey              string                    `json:"auth_key"`
	PasswordHash         string                    `json:"password_hash"`
	PasswordResetToken   string                    `json:"password_reset_token"`
	Email                string                    `json:"email"`
	Role                 int                       `json:"role"`
	Status               int8                      `json:"status"`
	CreateTime           string                    `json:"create_time"`
	UpdateTime           string                    `json:"update_time"`
	Goals                []*Goal                   `orm:"reverse(many)"`
	Conclusions          map[string]*Conclusion    `orm:"-"`
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(User))
}

func FindUser(id int) (user *User, err error) {
	o := orm.NewOrm()
	o.Using("default")

	user = &User{Id:id}
	err = o.Read(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this *User) Get(id int) (result map[string]interface{}, err error) {
	o := orm.NewOrm()
	o.Using("default")

	user, err := FindUser(id)
	if err != nil {
		return nil, err
	}

	o.LoadRelated(user, "Goals")

	result = map[string]interface{}{
		"email":user.Email,
	}
	return result, nil
}

func (this *User) GetInitialData() (result map[string]interface{}, err error) {
	o := orm.NewOrm()
	o.Using("default")

	o.LoadRelated(this, "Goals")

	days := []string{"today", "yesterday"}

	for _, goal := range this.Goals {
		goal.Today = &Day{}
		for _, day := range days {
			var day1 time.Time
			if day == "today" {
				day1 = time.Now()
			} else {
				day1 = time.Now().AddDate(0, 0, -1)
			}
			day2 := day1.AddDate(0, 0, +1)

			qs := o.QueryTable("report")
			qs = qs.Filter("report_date__gte", day1.Format("2014-07-08"))
			qs = qs.Filter("report_date__lt", day2.Format("2014-07-08"))

			if day == "today" {
				qs.All(&goal.Today.Reports)
			} else {
				qs.All(&goal.Yesterday.Reports)
			}
			//TODO: create if not found
		}
	}

	this.Conclusions = make(map[string]*Conclusion, len(days))

	for _, day := range days {
		this.Conclusions[day] = &Conclusion{}

		var day1 time.Time
		if day == "today" {
			day1 = time.Now()
		} else {
			day1 = time.Now().AddDate(0, 0, -1)
		}
		day2 := day1.AddDate(0, 0, +1)

		qs := o.QueryTable("conclusion")
		qs = qs.Filter("report_date__gte", day1.Format("2014-07-08"))
		qs = qs.Filter("report_date__lt", day2.Format("2014-07-08"))
		qs.All(this.Conclusions[day])

		//TODO: create if not found
	}

	var cats []*GoalCategory
	o.QueryTable("goal_category").All(&cats)

	catsReturn := make(map[string]*GoalCategory, len(cats))
	for _, cat := range cats {
		catsReturn[string(cat.Id)] = cat
	}

	result = map[string]interface{}{
		"categories": catsReturn,
		"goals": this.Goals,
		"conclusions": this.Conclusions,
	}

	return result, nil
}

