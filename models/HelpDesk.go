package models

import "time"

// HelpDesk represents a help desk ticket.
type HelpDesk struct {
	ID             int                   `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name           string                `gorm:"column:name;type:char;not null" json:"name"`
	CreateDate     time.Time             `gorm:"column:create_date;type:date;not null" json:"create_date"`
	HelpDeskTypeID int                   `gorm:"column:help_desk_type_id;not null" json:"help_desk_type_id"`
	PortalUserID   int                   `gorm:"column:portal_user_id;not null" json:"portal_user_id"`
	Description    string                `gorm:"column:description;type:text;not null" json:"description"`
	State          string                `gorm:"column:state;type:char;not null" json:"state"`
	ResolveDate    time.Time             `gorm:"column:resolve_date;type:date;not null" json:"resolve_date"`
	ParentID       *int                  `gorm:"column:help_desk_id" json:"parent_id,omitempty"`
	ProjectID      *int                  `gorm:"column:project_id" json:"project_id,omitempty"`
	Type           HelpDeskType          `gorm:"foreignKey:HelpDeskTypeID;references:ID" json:"type,omitempty"`
	PortalUser     User                  `gorm:"foreignKey:PortalUserID;references:ID" json:"portal_user,omitempty"`
	Parent         *HelpDesk             `gorm:"foreignKey:ParentID;references:ID" json:"parent,omitempty"`
	Documents      []Document            `gorm:"foreignKey:HelpDeskID;references:ID" json:"documents,omitempty"`
	Equipments     []Equipment           `gorm:"many2many:equipment_help_desk;foreignKey:ID;joinForeignKey:HelpDeskID;References:ID;joinReferences:EquipmentID" json:"equipments,omitempty"`
	Participants   []User                `gorm:"many2many:user_help_desk;foreignKey:ID;joinForeignKey:HelpDeskID;References:ID;joinReferences:UserID" json:"participants,omitempty"`
	Transactions   []HelpDeskTransaction `gorm:"foreignKey:HelpDeskID;references:ID" json:"transactions,omitempty"`
	Ratings        []HelpDeskRating      `gorm:"foreignKey:HelpDeskID;references:ID" json:"ratings,omitempty"`
}

// TableName overrides the default.
func (HelpDesk) TableName() string {
	return `public."Help_desk"`
}
