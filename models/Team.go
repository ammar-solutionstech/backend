package models

import "time"

// Team mirrors Help_desk_team records.
type Team struct {
	ID          int        `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"column:name;type:char;not null" json:"name"`
	ManagerID   *int       `gorm:"column:user_manager_id" json:"manager_id,omitempty"`
	Description *string    `gorm:"column:description;type:text" json:"description,omitempty"`
	IsActive    *bool      `gorm:"column:is_active" json:"is_active,omitempty"`
	CreatedAt   *time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"deleted_at,omitempty"`
	Manager     *User      `gorm:"foreignKey:ManagerID;references:ID" json:"manager,omitempty"`
	Users       []User     `gorm:"many2many:team_members;foreignKey:ID;joinForeignKey:HelpDeskTeamID;References:ID;joinReferences:UserID" json:"users,omitempty"`
}

// TableName ensures GORM queries the correct table in the public schema.
func (Team) TableName() string {
	return `public."Help_desk_team"`
}
