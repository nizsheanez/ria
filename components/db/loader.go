package db

import (
	"github.com/lann/squirrel"
	"database/sql"
	"ria/components"
)

type Loader struct {}

func NewLoader() *Loader {
	return &Loader{};
}

func (*Loader) FindByIdInTable(tableName string, id int, buf interface {}) error {
	qb := squirrel.Select("*").
	From(tableName).
	Where("id = ?", id)

	err := LoadOne(&qb, buf)
	if err != nil {
		return err
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

	err = components.App.Db.Unsafe().Select(buf, query, args...)
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

	err = components.App.Db.Unsafe().Get(buf, query, args...)
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

	err = components.App.Db.Unsafe().QueryRow(query, args...).Scan(buf)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	return nil
}
