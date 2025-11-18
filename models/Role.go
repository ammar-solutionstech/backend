package models

type Role struct {
	ID          int          `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name        string       `gorm:"column:name;type:varchar(255);not null" json:"name"`
	Description *string      `gorm:"column:description;type:text" json:"description,omitempty"`
	Permissions []Permission `gorm:"many2many:role_permission;joinForeignKey:RoleID;joinReferences:PermissionID" json:"permissions,omitempty"`
	Users       []User       `gorm:"many2many:user_role;joinForeignKey:RoleID;joinReferences:UserID" json:"users,omitempty"`
}

func (Role) TableName() string {
	return `public."Role"`
}

type RolePermission struct {
	RoleID       int `gorm:"column:role_id" json:"role_id"`
	PermissionID int `gorm:"column:permission_id" json:"permission_id"`
}

func (RolePermission) TableName() string {
	return `public."role_permission"`
}
