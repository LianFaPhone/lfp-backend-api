package models

type Role struct {
	Id     int64  `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Status int64  `json:"status" gorm:"column:status"`
	Label  string `json:"label" gorm:"column:label"`
	Model
}

func (c *Role) TableName() string {
	return "t_roles"
}
