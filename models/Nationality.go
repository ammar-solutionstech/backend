package models

// Country represents entries in the public.Country table.
type Country struct {
	ID   int    `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Code string `gorm:"column:code;type:varchar(50);not null" json:"code"`
}

// TableName overrides the schema-qualified table name for Country.
func (Country) TableName() string {
	return `public."Country"`
}

// Nationality is kept as an alias for backwards compatibility with existing code.
type Nationality = Country
