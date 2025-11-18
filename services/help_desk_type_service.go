package services

import (
	"errors"

	"gorm.io/gorm"

	"backend/models"
)

// HelpDeskTypeService contains the data access logic for help desk types.
type HelpDeskTypeService struct {
	db *gorm.DB
}

// NewHelpDeskTypeService constructs a new HelpDeskTypeService.
func NewHelpDeskTypeService(db *gorm.DB) *HelpDeskTypeService {
	return &HelpDeskTypeService{db: db}
}

// List retrieves all help desk types.
func (s *HelpDeskTypeService) List() ([]models.HelpDeskType, error) {
	var types []models.HelpDeskType
	if err := s.db.Order("id ASC").Find(&types).Error; err != nil {
		return nil, err
	}
	return types, nil
}

// Get retrieves a single help desk type by its identifier.
func (s *HelpDeskTypeService) Get(id int) (*models.HelpDeskType, error) {
	var helpDeskType models.HelpDeskType
	if err := s.db.First(&helpDeskType, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &helpDeskType, nil
}

// Create persists a new help desk type.
func (s *HelpDeskTypeService) Create(helpDeskType *models.HelpDeskType) error {
	return s.db.Create(helpDeskType).Error
}

// Update updates an existing help desk type.
func (s *HelpDeskTypeService) Update(id int, updates map[string]interface{}) (*models.HelpDeskType, error) {
	helpDeskType := &models.HelpDeskType{}
	if err := s.db.First(helpDeskType, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	if err := s.db.Model(helpDeskType).Updates(updates).Error; err != nil {
		return nil, err
	}

	return helpDeskType, nil
}

// Delete removes a help desk type by ID.
func (s *HelpDeskTypeService) Delete(id int) error {
	result := s.db.Delete(&models.HelpDeskType{}, id)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
