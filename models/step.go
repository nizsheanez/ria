package models

type Step struct {
	Id                      int            `db:"id" json:"id"`
	Title                   string    	   `db:"title" json:"title"`
	Status                  int8           `db:"report_date" json:"report_date"`
	CreateTime              string         `db:"create_time" json:"create_time"`
	UpdateTime              string         `db:"update_time" json:"update_time"`
}

func init() {
	// Need to register model in init
//	orm.RegisterModel(new(Step))
}
