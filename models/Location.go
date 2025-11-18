package models

// Location represents a physical location for equipment.
type Location struct {
	ID             int      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	Name           string   `gorm:"column:name;type:char;not null" json:"name"`
	ZipCode        *string  `gorm:"column:zip_code;type:char;not null" json:"zip_code"`
	State          *string  `gorm:"column:state;type:char;not null" json:"state"`
	BuildingNumber *int     `gorm:"column:building_number;not null" json:"building_number"`
	RoomNumber     *int     `gorm:"column:room_number;not null" json:"room_number"`
	Latitude       *float64 `gorm:"column:latitude;type:numeric;not null" json:"latitude"`
	Longitude      *float64 `gorm:"column:longitude;type:numeric;not null" json:"longitude"`
	CountryID      *int     `gorm:"column:country_id;not null" json:"country_id"`
	CityID         *int     `gorm:"column:city_id;not null" json:"city_id"`
	LocationMap    *string  `gorm:"column:location_map;type:text;not null" json:"location_map"`
	Country        *Country `gorm:"foreignKey:CountryID;references:ID" json:"country,omitempty"`
	City           *City    `gorm:"foreignKey:CityID;references:ID" json:"city,omitempty"`
}

// TableName overrides the default.
func (Location) TableName() string {
	return `public."Location"`
}
