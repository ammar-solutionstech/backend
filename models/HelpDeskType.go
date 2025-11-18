package models

import "time"

// HelpDeskType represents a type/category for help desk tickets.
// The struct tags map the Go fields to the existing database columns.
type HelpDeskType struct {
	ID          int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description string     `gorm:"column:description;type:text;not null" json:"description"`
	TeamID      *int       `gorm:"column:team_id" json:"team_id,omitempty"`
	IsActive    *bool      `gorm:"column:is_active" json:"is_active,omitempty"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Team        *Team      `gorm:"foreignKey:TeamID;references:ID" json:"team,omitempty"`
}

// TableName ensures GORM queries the correct table in the public schema.
func (HelpDeskType) TableName() string {
	return `public."Help_desk_type"`
}
