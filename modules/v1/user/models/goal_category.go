package models

type GoalCategory struct {
	Id                 int              `db:"id" json:"id"`
	Name               string           `db:"name" json:"name"`
	CreateTime         string           `db:"create_time" json:"create_time"`
	UpdateTime         string           `db:"update_time" json:"update_time"`
}

func init() {
	// Need to register model in init
//	orm.RegisterModel(new(GoalCategory))
}

