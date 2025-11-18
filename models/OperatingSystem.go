package models

// OperatingSystem represents OS metadata for equipment records.
type OperatingSystem struct {
	ID            int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"column:name;type:char;not null" json:"name"`
	Description   string `gorm:"column:description;type:char;not null" json:"description"`
	Version       string `gorm:"column:version;type:char;not null" json:"version"`
	Architectures string `gorm:"column:architectures;type:char;not null" json:"architectures"`
}

// TableName overrides the default.
func (OperatingSystem) TableName() string {
	return `public."Operating_System"`
}
