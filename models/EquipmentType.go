package models

// EquipmentType categorizes equipment (e.g., laptop, server).
type EquipmentType struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"column:name;type:char;not null" json:"name"`
	Description string `gorm:"column:description;type:char;not null" json:"description"`
}

// TableName overrides the default table name.
func (EquipmentType) TableName() string {
	return `public."Equipment_type"`
}
