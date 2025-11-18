package models

// Document stores binary documents linked to equipment, suppliers, or help desk tickets.
type Document struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name         string    `gorm:"column:name;type:char;not null" json:"name"`
	DocumentSize string    `gorm:"column:document_size;type:char;not null" json:"document_size"`
	Picture      []byte    `gorm:"column:picture;type:bytea;not null" json:"picture"`
	DocumentType string    `gorm:"column:document_type;type:char;not null" json:"document_type"`
	EquipmentID  int       `gorm:"column:equipment_id;not null" json:"equipment_id"`
	SupplierID   int       `gorm:"column:supplier_id;not null" json:"supplier_id"`
	HelpDeskID   int       `gorm:"column:help_desk_id;not null" json:"help_desk_id"`
	Equipment    Equipment `gorm:"foreignKey:EquipmentID;references:ID" json:"equipment,omitempty"`
	Supplier     Contact   `gorm:"foreignKey:SupplierID;references:ID" json:"supplier,omitempty"`
	HelpDesk     HelpDesk  `gorm:"foreignKey:HelpDeskID;references:ID" json:"help_desk,omitempty"`
}

// TableName overrides the default table name.
func (Document) TableName() string {
	return `public."Documents"`
}
