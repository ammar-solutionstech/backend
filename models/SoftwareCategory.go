package models

// SoftwareCategory represents categories for software records.
type SoftwareCategory struct {
	ID          int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string `gorm:"column:name;type:char;not null" json:"name"`
	Description string `gorm:"column:description;type:char;not null" json:"description"`
}

// TableName overrides default.
func (SoftwareCategory) TableName() string {
	return `public."Software_category"`
}
