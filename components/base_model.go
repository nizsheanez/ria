package components

import (
	"github.com/lann/squirrel"
	"database/sql"
	"github.com/astaxie/beego/validation"
)

type BaseModel struct {

}


func (this *BaseModel) Validate(scenario string) (b bool, err error) {
	validator := validation.Validation{}

	if this, ok :=  interface{}(this).(Validatable); ok {
		this.Validators(&validator, scenario)
	} else {
		panic("Model don't implement interface Validatable")
	}

	return validator.Valid(this)
}

type Validatable interface {
	Validators(validator *validation.Validation, scenario string)
}

type TableMapper interface {
	TableName() string
	FindById(id int) error
}

func (this *BaseModel) FindById(id int) error {
	if this, ok :=  interface{}(this).(TableMapper); ok {
		qb := squirrel.Select("*").
			From(this.TableName()).
			Where("id = ?", id)

		err := LoadOne(&qb, this)
		if err != nil {
			return err
		}
	} else {
		panic("Model don't implement interface Validatable")
	}

	return nil
}


// loadCollection uses QueryBuilder to run SELECT query and fetch collection of records from database
// *search* is used to specify LIMIT and OFFSET for SELECT query
// *search* can be nil. In this case it won't be used
func LoadCollection(qb *squirrel.SelectBuilder, buf interface{}) error {

	query, args, err := qb.ToSql()

	if err != nil {
		return err
	}

	err = App.Db.Unsafe().Select(buf, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func LoadOne(qb *squirrel.SelectBuilder, buf interface{}) error {

	query, args, err := qb.ToSql()

	if err != nil {
		return err
	}

	err = App.Db.Unsafe().Get(buf, query, args...)
	if err != nil {
		return err
	}

	return nil
}

// loadValue uses QueryBuilder to run SELECT query and fetch single value from database,
// use it for getting ID or counter
func LoadValue(qb *squirrel.SelectBuilder, buf interface{}) error {
	query, args, err := qb.ToSql()

	if err != nil {
		return err
	}

	err = App.Db.Unsafe().QueryRow(query, args...).Scan(buf)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}
