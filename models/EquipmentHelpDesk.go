package models

// EquipmentHelpDesk links equipment to help desk tickets.
type EquipmentHelpDesk struct {
	EquipmentID int `gorm:"column:equipment_id" json:"equipment_id"`
	HelpDeskID  int `gorm:"column:help_desk_id" json:"help_desk_id"`
}

// TableName overrides default.
func (EquipmentHelpDesk) TableName() string {
	return `public."equipment_help_desk"`
}
