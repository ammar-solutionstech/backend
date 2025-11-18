package models

// TransactionType defines help desk transaction categories.
type TransactionType struct {
	ID           int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string  `gorm:"column:name;type:char;not null" json:"name"`
	ExpectedTime float64 `gorm:"column:expected_time;not null" json:"expected_time"`
}

// TableName overrides default.
func (TransactionType) TableName() string {
	return `public."Transaction_type"`
}
