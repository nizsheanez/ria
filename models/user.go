package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"strconv"
	"fmt"
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
	Goals                []*Goal                   `orm:"reverse(many)" json:"-"`
	Conclusions          map[string]*Conclusion    `orm:"-" json:"-"`
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

	type ResponseGoal struct {
		Goal
		Today              *Day                    `json:"today"`
		Yesterday          *Day                    `json:"yesterday"`
	}

	type Response struct {
		Categories  map[string]*GoalCategory `json:"categories"`
		Goals       []*ResponseGoal `json:"goals"`
		Conclusions map[string]*Conclusion `json:"conclusions"`
	}

	response := &Response{}

	var rawGoals []*Goal
	qs := o.QueryTable("goal")
	qs.Filter("fk_user", this.Id)
	qs.All(&rawGoals)

	days := []string{"today", "yesterday"}

	response.Goals = make([]*ResponseGoal, len(rawGoals))
	for i, goal := range rawGoals {
		response.Goals[i] = &ResponseGoal{Goal: *goal}

		response.Goals[i].Today = &Day{Report: &Report{}}
		response.Goals[i].Yesterday = &Day{Report: &Report{}}
		for _, day := range days {
			var day1 time.Time
			if day == "today" {
				day1 = time.Now()
			} else {
				day1 = time.Now().AddDate(0, 0, -1)
			}
			day2 := day1.AddDate(0, 0, +1)

			qs := o.QueryTable("report")
			qs = qs.Filter("report_date__gte", FormatDate(day1))
			qs = qs.Filter("report_date__lt", FormatDate(day2))
			qs = qs.Filter("fk_goal", response.Goals[i].Id)

			if day == "today" {
				var r Report
				qs.One(&r)
				response.Goals[i].Today.Report = &r
				beego.Info(fmt.Sprintf("%v - %v", response.Goals[i].Yesterday.Report, r))
				panic(1)
			} else {
				var r Report
				qs.One(response.Goals[i].Yesterday.Report)
				response.Goals[i].Yesterday.Report = &r
				beego.Info(fmt.Sprintf("%v - %v", response.Goals[i].Yesterday.Report, r))
				panic(1)
			}
			//TODO: create if not found
		}
	}

	response.Conclusions = make(map[string]*Conclusion, len(days))

	for _, day := range days {
		response.Conclusions[day] = &Conclusion{}

		var day1 time.Time
		if day == "today" {
			day1 = time.Now()
		} else {
			day1 = time.Now().AddDate(0, 0, -1)
		}
		day2 := day1.AddDate(0, 0, +1)

		qs := o.QueryTable("conclusion")
		qs = qs.Filter("report_date__gte", FormatDate(day1))
		qs = qs.Filter("report_date__lt", FormatDate(day2))
		qs.All(response.Conclusions[day])

		//TODO: create if not found
	}

	var cats []*GoalCategory
	o.QueryTable("goal_category").All(&cats)

	response.Categories = make(map[string]*GoalCategory, len(cats))
	for _, cat := range cats {
		response.Categories[strconv.Itoa(cat.Id)] = cat
	}

	result = map[string]interface{}{
		"categories": response.Categories,
		"goals": response.Goals,
		"conclusions": response.Conclusions,
	}

	return result, nil
}

func FormatDate(t time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", t.Year(), t.Month(), t.Day())
}
