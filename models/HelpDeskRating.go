package models

// HelpDeskRating stores feedback for help desk tickets.
type HelpDeskRating struct {
	ID         int      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name       string   `gorm:"column:name;type:char;not null" json:"name"`
	Rate       int      `gorm:"column:rate;not null" json:"rate"`
	HelpDeskID int      `gorm:"column:help_desk_id;not null" json:"help_desk_id"`
	HelpDesk   HelpDesk `gorm:"foreignKey:HelpDeskID;references:ID" json:"help_desk,omitempty"`
}

// TableName overrides default.
func (HelpDeskRating) TableName() string {
	return `public."Help_desk_rating"`
}
