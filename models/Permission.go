package models

type Permission struct {
	ID     int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Module string `gorm:"column:module;type:varchar(255);not null" json:"module"`
}

func (Permission) TableName() string {
	return `public."Permission"`
}
