package models

type Goal struct {
	Id                   int                      `db:"id" json:"id"`
	FkUser               int                      `db:"fk_user" json:"fk_user"`
	FkGoalCategory       int                      `db:"fk_goal_category" json:"fk_goal_category"`
	Title                NullString               `db:"title" json:"title"`
	Status               int                      `db:"status" json:"status"`
	Completed            int                      `db:"completed" json:"completed"`
	Reason               NullString               `db:"reason" json:"reason"`
	CreateTime           NullString               `db:"create_time" json:"create_time"`
	UpdateTime           NullString               `db:"update_time" json:"update_time"`
	Decomposition        NullString               `db:"decomposition" json:"decomposition"`
	Comments             NullString               `db:"comments" json:"comments"`
}

type Day struct {
	Report            *Report     `orm:"-" json:"report"`
}

func init() {
	// Need to register model in init
	//	orm.RegisterModel(new(Goal))
}

