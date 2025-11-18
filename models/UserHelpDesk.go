package models

// UserHelpDesk links users to help desk tickets.
type UserHelpDesk struct {
	UserID     int `gorm:"column:user_id" json:"user_id"`
	HelpDeskID int `gorm:"column:help_desk_id" json:"help_desk_id"`
}

// TableName overrides default.
func (UserHelpDesk) TableName() string {
	return `public."user_help_desk"`
}
