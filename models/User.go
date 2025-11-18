package models

import "time"

// User maps to the public.User table and carries identity/contact data.
type User struct {
	ID              int                   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FirstName       string                `gorm:"column:first_name;type:varchar(255);not null" json:"first_name"`
	LatestName      string                `gorm:"column:latest_name;type:varchar(255);not null" json:"latest_name"`
	FatherName      string                `gorm:"column:father_name;type:varchar(255);not null" json:"father_name"`
	NationalityID   *int                  `gorm:"column:nationality_id" json:"nationality_id,omitempty"`
	DateOfBirth     *time.Time            `gorm:"column:dob;type:date" json:"dob,omitempty"`
	WorkEmail       string                `gorm:"column:work_email;type:varchar(255);not null" json:"work_email"`
	PrivateEmail    *string               `gorm:"column:private_email;type:varchar(255)" json:"private_email,omitempty"`
	Password        string                `gorm:"column:password;type:varchar(255);not null" json:"password,omitempty"`
	WorkMobile      *string               `gorm:"column:work_mobile;type:varchar(255)" json:"work_mobile,omitempty"`
	PrivateMobile   *string               `gorm:"column:private_mobile;type:varchar(255)" json:"private_mobile,omitempty"`
	IDNumber        *string               `gorm:"column:id_number;type:varchar(255)" json:"id_number,omitempty"`
	IDType          *int                  `gorm:"column:id_type" json:"id_type,omitempty"`
	ContactID       *int                  `gorm:"column:contact_id" json:"contact_id,omitempty"`
	Active          *int                  `gorm:"column:active" json:"active,omitempty"`
	Roles           []Role                `gorm:"many2many:user_role;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:RoleID" json:"roles,omitempty"`
	Permissions     []Permission          `gorm:"many2many:user_special_permission;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:PermissionID" json:"permissions,omitempty"`
	Teams           []Team                `gorm:"many2many:team_members;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:HelpDeskTeamID" json:"teams,omitempty"`
	Nationality     *Country              `gorm:"foreignKey:NationalityID;references:ID" json:"nationality,omitempty"`
	Contact         *Contact              `gorm:"foreignKey:ContactID;references:ID" json:"contact,omitempty"`
	HelpDeskTickets []HelpDesk            `gorm:"many2many:user_help_desk;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:HelpDeskID" json:"help_desk_tickets,omitempty"`
	Transactions    []HelpDeskTransaction `gorm:"many2many:Help_desk_transaction_user;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:TransactionID" json:"transactions,omitempty"`
}

// TableName overrides the schema-qualified table name.
func (User) TableName() string {
	return `public."User"`
}

// UserPermission represents the user_special_permission join table.
type UserPermission struct {
	UserID       int `gorm:"column:user_id" json:"user_id"`
	PermissionID int `gorm:"column:permission_id" json:"permission_id"`
}

// TableName overrides the schema-qualified table name.
func (UserPermission) TableName() string {
	return `public."user_special_permission"`
}

// UserRole represents the user_role join table.
type UserRole struct {
	UserID int `gorm:"column:user_id" json:"user_id"`
	RoleID int `gorm:"column:role_id" json:"role_id"`
}

// TableName overrides the schema-qualified table name.
func (UserRole) TableName() string {
	return `public."user_role"`
}

// TeamMember represents the team_members join table.
type TeamMember struct {
	UserID         int `gorm:"column:user_id" json:"user_id"`
	HelpDeskTeamID int `gorm:"column:help_desk_team_id" json:"help_desk_team_id"`
}

// TableName overrides the schema-qualified table name.
func (TeamMember) TableName() string {
	return `public."team_members"`
}
