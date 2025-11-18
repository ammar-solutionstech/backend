package controllers

import (
	"gorm.io/gorm"

	"backend/models"
)

func NewBrandResource(db *gorm.DB) *ResourceController[models.Brand] {
	return NewResourceController[models.Brand](db, "brand", []string{"name"})
}

func NewCityResource(db *gorm.DB) *ResourceController[models.City] {
	return NewResourceController[models.City](db, "city", []string{"name", "country_id"})
}

func NewContactResource(db *gorm.DB) *ResourceController[models.Contact] {
	return NewResourceController[models.Contact](db, "contact", []string{
		"name", "country_id", "mobile_number", "phone_number", "website", "email",
	})
}

func NewCountryResource(db *gorm.DB) *ResourceController[models.Country] {
	return NewResourceController[models.Country](db, "country", []string{"name", "code"})
}

func NewDocumentResource(db *gorm.DB) *ResourceController[models.Document] {
	return NewResourceController[models.Document](db, "document", []string{
		"name", "document_size", "picture", "document_type", "equipment_id", "supplier_id", "help_desk_id",
	})
}

func NewEquipmentResource(db *gorm.DB) *ResourceController[models.Equipment] {
	return NewResourceController[models.Equipment](db, "equipment", []string{
		"name", "description", "state", "serial_number", "equipment_type_id", "model_id",
		"production_date", "ip_v4", "ip_v6", "mac_address", "operating_system_id", "location_id",
		"proccessors", "country_of_region", "user_id", "warrantly_start_date", "warrantly_end_date",
		"supplier_id",
	})
}

func NewEquipmentTypeResource(db *gorm.DB) *ResourceController[models.EquipmentType] {
	return NewResourceController[models.EquipmentType](db, "equipment_type", []string{"name", "description"})
}

func NewHelpDeskResource(db *gorm.DB) *ResourceController[models.HelpDesk] {
	return NewResourceController[models.HelpDesk](db, "help_desk", []string{
		"name", "create_date", "help_desk_type_id", "portal_user_id", "description", "state",
		"resolve_date", "parent_id", "project_id",
	})
}

func NewHelpDeskRatingResource(db *gorm.DB) *ResourceController[models.HelpDeskRating] {
	return NewResourceController[models.HelpDeskRating](db, "help_desk_rating", []string{
		"name", "rate", "help_desk_id",
	})
}

func NewTeamResource(db *gorm.DB) *ResourceController[models.Team] {
	return NewResourceController[models.Team](db, "team", []string{
		"name", "user_manager_id", "description", "is_active", "created_at", "updated_at", "deleted_at",
	})
}

func NewHelpDeskTransactionResource(db *gorm.DB) *ResourceController[models.HelpDeskTransaction] {
	return NewResourceController[models.HelpDeskTransaction](db, "help_desk_transaction", []string{
		"name", "transaction_type_id", "date_time", "time", "help_desk_id",
	})
}

func NewLocationResource(db *gorm.DB) *ResourceController[models.Location] {
	return NewResourceController[models.Location](db, "location", []string{
		"name", "zip_code", "state", "building_number", "room_number", "latitude", "longitude",
		"country_id", "city_id", "location_map",
	})
}

func NewMaintenanceResource(db *gorm.DB) *ResourceController[models.Maintenance] {
	return NewResourceController[models.Maintenance](db, "maintenance", []string{
		"name", "start_date", "end_date", "contact_id",
	})
}

func NewMenuResource(db *gorm.DB) *ResourceController[models.Menu] {
	return NewResourceController[models.Menu](db, "menu", []string{
		"name", "title", "uri",
	})
}

func NewModelResource(db *gorm.DB) *ResourceController[models.Model] {
	return NewResourceController[models.Model](db, "model", []string{
		"name", "brand_id",
	})
}

func NewOperatingSystemResource(db *gorm.DB) *ResourceController[models.OperatingSystem] {
	return NewResourceController[models.OperatingSystem](db, "operating_system", []string{
		"name", "description", "version", "architectures",
	})
}

func NewSoftwareResource(db *gorm.DB) *ResourceController[models.Software] {
	return NewResourceController[models.Software](db, "software", []string{
		"software_name", "category_id", "license_exp_date",
	})
}

func NewSoftwareCategoryResource(db *gorm.DB) *ResourceController[models.SoftwareCategory] {
	return NewResourceController[models.SoftwareCategory](db, "software_category", []string{
		"name", "description",
	})
}

func NewTransactionTypeResource(db *gorm.DB) *ResourceController[models.TransactionType] {
	return NewResourceController[models.TransactionType](db, "transaction_type", []string{
		"name", "expected_time",
	})
}
