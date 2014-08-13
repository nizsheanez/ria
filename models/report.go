package models

type Report struct {
	Id                      int            		`db:"id" json:"id"`
	Description             NullString     		`db:"description" json:"description"`
	ReportDate              NullString     		`db:"report_date" json:"report_date"`
	CreateTime              string         		`db:"create_time" json:"create_time"`
	UpdateTime              string         		`db:"update_time" json:"update_time"`
}

func init() {
	// Need to register model in init
//	orm.RegisterModel(new(Report))
}
