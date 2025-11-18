package models

// MenuRole represents role assignments to menus.
type MenuRole struct {
	MenuID int `gorm:"column:menu_id" json:"menu_id"`
	RoleID int `gorm:"column:role_id" json:"role_id"`
}

// TableName overrides default.
func (MenuRole) TableName() string {
	return `public."menu_role"`
}
