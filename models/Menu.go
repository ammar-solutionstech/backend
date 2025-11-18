package models

// Menu represents navigation items that can be assigned to roles.
type Menu struct {
	ID    int     `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name  string  `gorm:"column:name;type:char;not null" json:"name"`
	Title *string `gorm:"column:title;type:char" json:"title,omitempty"`
	URI   *string `gorm:"column:uri;type:char" json:"uri,omitempty"`
	Roles []Role  `gorm:"many2many:menu_role;foreignKey:ID;joinForeignKey:MenuID;References:ID;joinReferences:RoleID" json:"roles,omitempty"`
}

// TableName overrides default.
func (Menu) TableName() string {
	return `public."Menu"`
}
