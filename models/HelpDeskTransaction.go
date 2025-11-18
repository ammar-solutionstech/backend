package models

import "time"

// HelpDeskTransaction captures state transitions for help desk tickets.
type HelpDeskTransaction struct {
	ID                int             `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name              string          `gorm:"column:name;type:char;not null" json:"name"`
	TransactionTypeID int             `gorm:"column:transaction_type_id;not null" json:"transaction_type_id"`
	DateTime          time.Time       `gorm:"column:date_time;type:time;not null" json:"date_time"`
	TimeSpent         float64         `gorm:"column:time;not null" json:"time"`
	HelpDeskID        int             `gorm:"column:help_desk_id;not null" json:"help_desk_id"`
	Type              TransactionType `gorm:"foreignKey:TransactionTypeID;references:ID" json:"type,omitempty"`
	HelpDesk          HelpDesk        `gorm:"foreignKey:HelpDeskID;references:ID" json:"help_desk,omitempty"`
	Users             []User          `gorm:"many2many:Help_desk_transaction_user;foreignKey:ID;joinForeignKey:TransactionID;References:ID;joinReferences:UserID" json:"users,omitempty"`
}

// TableName overrides default.
func (HelpDeskTransaction) TableName() string {
	return `public."Help_desk_transaction"`
}
