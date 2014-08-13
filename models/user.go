package models

import (
	"time"
	"github.com/astaxie/beego/orm"
//		"github.com/astaxie/beego"
	"ria/components"
	"database/sql"
	"github.com/lann/squirrel"
	"strconv"
	"fmt"
)

type User struct {
	Id                   int                       `db:"id" json:"id"`
	UserName             NullString                `db:"username" json:"username"`
	AuthKey              NullString                `db:"auth_key" json:"auth_key"`
	PasswordHash         NullString                `db:"password_hash" json:"password_hash"`
	PasswordResetToken   NullString                `db:"password_reset_token" json:"password_reset_token"`
	Email                NullString                `db:"email" json:"email"`
	Role                 int                       `db:"role" json:"role"`
	Status               int8                      `db:"status" json:"status"`
	CreateTime           string                    `db:"create_time" json:"create_time"`
	UpdateTime           string                    `db:"update_time" json:"update_time"`
	//	Goals                []*Goal                   `orm:"reverse(many)" json:"-"`
	//	Conclusions          map[string]*Conclusion    `orm:"-" json:"-"`
}

func init() {
	// Need to register model in init
	//	orm.RegisterModel(new(User))
}

func FindUser(id int) (user *User, err error) {
	user = &User{}
	qb := squirrel.Select("*").
	From("user").
	Where("id = ?", id)

	err = loadOne(&qb, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this *User) Get(id int) (result map[string]interface{}, err error) {

	user := &User{}
	qb := squirrel.Select("*").
	From("user").
	Where("id = ?", id)
	err = loadOne(&qb, user)
	if err != nil {
		return nil, err
	}

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
	qb := squirrel.Select("*").
	From("goal").
	Where("fk_user = ?", this.Id)

	err = loadCollection(&qb, &rawGoals)
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
				err = loadOne(&qb, response.Goals[i].Today.Report)
				if err != nil {
					return nil, err
				}
				//TODO: create if not found
			} else {
				err = loadOne(&qb, response.Goals[i].Yesterday.Report)
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

		err = loadOne(&qb, response.Conclusions[day])
		if err != nil {
			return nil, err
		}

		//TODO: create if not found
	}

	var cats []*GoalCategory
	qb = squirrel.Select("*").
	From("goal_category")

	err = loadCollection(&qb, &cats)
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

// loadCollection uses QueryBuilder to run SELECT query and fetch collection of records from database
// *search* is used to specify LIMIT and OFFSET for SELECT query
// *search* can be nil. In this case it won't be used
func loadCollection(qb *squirrel.SelectBuilder, buf interface{}) error {

	query, args, err := qb.ToSql()

	if err != nil {
		return err
	}

	err = components.App.Db.Unsafe().Select(buf, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func loadOne(qb *squirrel.SelectBuilder, buf interface{}) error {

	query, args, err := qb.ToSql()

	if err != nil {
		return err
	}

	err = components.App.Db.Unsafe().Get(buf, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// loadValue uses QueryBuilder to run SELECT query and fetch single value from database,
// use it for getting ID or counter
func loadValue(qb *squirrel.SelectBuilder, buf interface{}) error {
	query, args, err := qb.ToSql()

	if err != nil {
		return err
	}

	err = components.App.Db.Unsafe().QueryRow(query, args...).Scan(buf)
	if err != nil && err != sql.ErrNoRows {
		err = err
	}

	return err
}
