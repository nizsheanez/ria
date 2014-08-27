package models

type Goal struct {
	Id                   int                      `db:"id" json:"id"`
	FkUser               int                      `db:"fk_user" json:"fk_user"`
	FkGoalCategory       int                      `db:"fk_goal_category" json:"fk_goal_category"`
	Title                string               	  `db:"title" json:"title"`
	Status               int                      `db:"status" json:"status"`
	Completed            int                      `db:"completed" json:"completed"`
	Reason               string               	  `db:"reason" json:"reason"`
	CreateTime           string               	  `db:"create_time" json:"create_time"`
	UpdateTime           string               	  `db:"update_time" json:"update_time"`
	Decomposition        string                   `db:"decomposition" json:"decomposition"`
	Comments             string               	  `db:"comments" json:"comments"`
}

type Day struct {
	Report            *Report     `orm:"-" json:"report"`
}

func init() {
	// Need to register model in init
	//	orm.RegisterModel(new(Goal))
}

