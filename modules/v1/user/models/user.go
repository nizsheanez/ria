package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	//		"github.com/astaxie/beego"
	"ria/components"
	"github.com/astaxie/beego/validation"
	"github.com/lann/squirrel"
	"strconv"
	"fmt"
	"ria/components/db"

)

type User struct {
	components.BaseModel
	Id                   int                   `db:"id" json:"id"`
	UserName             string                `db:"username" json:"username"`
	AuthKey              string               `db:"auth_key" json:"auth_key"`
	PasswordHash         string                `db:"password_hash" json:"password_hash"`
	PasswordResetToken   string                `db:"password_reset_token" json:"password_reset_token"`
	Email                string                `db:"email" json:"email"`
	Role                 int                   `db:"role" json:"role"`
	Status               int8                  `db:"status" json:"status"`
	CreateTime           string                `db:"create_time" json:"create_time"`
	UpdateTime           string                `db:"update_time" json:"update_time"`
	//	Goals                []*Goal                   `orm:"reverse(many)" json:"-"`
	//	Conclusions          map[string]*Conclusion    `orm:"-" json:"-"`
}

func init() {
	// Need to register model in init
	//	orm.RegisterModel(new(User))
}

func NewUser() *User {
	return &User{}
}

func (*User) TableName() string {
	return "user"
}

func FindUsers() ([]User, error) {
	qb := squirrel.Select("*").From("user")

	var users []User
	err := db.LoadCollection(&qb, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (this *User) Get(id int) (result map[string]interface{}, err error) {

	user := &User{}
	qb := squirrel.Select("*").
	From("user").
	Where("id = ?", id)
	err = db.LoadOne(&qb, user)
	if err != nil {
		return nil, err
	}

	result = map[string]interface{}{
		"email":user.Email,
	}
	return result, nil
}

func (this *User) FindById(id int) error {
	err := db.NewLoader().FindByIdInTable(this.TableName(), id, this)
	if err != nil {
		return err
	}

	return nil
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
	qb := squirrel.Select("*").
	From("goal").
	Where("fk_user = ?", this.Id)

	err = db.LoadCollection(&qb, &rawGoals)
	if err != nil {
		return nil, err
	}

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

			qb = squirrel.Select("*").
			From("report").
			Where("report_date >= ?", FormatDate(day1)).
			Where("report_date < ?", FormatDate(day2)).
			Where("fk_goal = ?", response.Goals[i].Id)

			if day == "today" {
				err = db.LoadOne(&qb, response.Goals[i].Today.Report)
				if err != nil {
					return nil, err
				}
				//TODO: create if not found
			} else {
				err = db.LoadOne(&qb, response.Goals[i].Yesterday.Report)
				if err != nil {
					return nil, err
				}
				//TODO: create if not found
			}
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

		qb = squirrel.Select("*").
		From("conclusion").
		Where("report_date >= ?", FormatDate(day1)).
		Where("report_date < ?", FormatDate(day2))

		err = db.LoadOne(&qb, response.Conclusions[day])
		if err != nil {
			return nil, err
		}

		//TODO: create if not found
	}

	var cats []*GoalCategory
	qb = squirrel.Select("*").
	From("goal_category")

	err = db.LoadCollection(&qb, &cats)
	if err != nil {
		return nil, err
	}
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

func (this *User) Validators(validator *validation.Validation, scenario string) {
}
