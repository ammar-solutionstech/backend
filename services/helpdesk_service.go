package services

import (
	"gorm.io/gorm"

	"backend/models"
)

// HelpDeskService encapsulates help desk relationship management.
type HelpDeskService struct {
	db *gorm.DB
}

// NewHelpDeskService constructs a HelpDeskService.
func NewHelpDeskService(db *gorm.DB) *HelpDeskService {
	return &HelpDeskService{db: db}
}

// Team membership helpers.

func (s *HelpDeskService) ListTeamMembers(teamID int) ([]models.TeamMember, error) {
	var members []models.TeamMember
	if err := s.db.Where("help_desk_team_id = ?", teamID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (s *HelpDeskService) AddTeamMember(teamID, userID int) error {
	member := models.TeamMember{
		HelpDeskTeamID: teamID,
		UserID:         userID,
	}
	return s.db.Where(member).FirstOrCreate(&member).Error
}

func (s *HelpDeskService) RemoveTeamMember(teamID, userID int) error {
	result := s.db.Where("help_desk_team_id = ? AND user_id = ?", teamID, userID).
		Delete(&models.TeamMember{})
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Help desk participant helpers.

func (s *HelpDeskService) ListParticipants(helpDeskID int) ([]models.UserHelpDesk, error) {
	var participants []models.UserHelpDesk
	if err := s.db.Where("help_desk_id = ?", helpDeskID).Find(&participants).Error; err != nil {
		return nil, err
	}
	return participants, nil
}

func (s *HelpDeskService) AddParticipant(helpDeskID, userID int) error {
	record := models.UserHelpDesk{
		HelpDeskID: helpDeskID,
		UserID:     userID,
	}
	return s.db.Where(record).FirstOrCreate(&record).Error
}

func (s *HelpDeskService) RemoveParticipant(helpDeskID, userID int) error {
	result := s.db.Where("help_desk_id = ? AND user_id = ?", helpDeskID, userID).
		Delete(&models.UserHelpDesk{})
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Help desk transaction assignment helpers.

func (s *HelpDeskService) ListTransactionUsers(transactionID int) ([]models.HelpDeskTransactionUser, error) {
	var records []models.HelpDeskTransactionUser
	if err := s.db.Where("transaction_id = ?", transactionID).Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (s *HelpDeskService) AddTransactionUser(transactionID, userID int) error {
	record := models.HelpDeskTransactionUser{
		TransactionID: transactionID,
		UserID:        userID,
	}
	return s.db.Where(record).FirstOrCreate(&record).Error
}

func (s *HelpDeskService) RemoveTransactionUser(transactionID, userID int) error {
	result := s.db.Where("transaction_id = ? AND user_id = ?", transactionID, userID).
		Delete(&models.HelpDeskTransactionUser{})
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// Validation helpers.

func (s *HelpDeskService) ValidateTeam(teamID int) error {
	var exists bool
	if err := s.db.Model(&models.Team{}).Select("count(1) > 0").Where("id = ?", teamID).Scan(&exists).Error; err != nil {
		return err
	}
	if !exists {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *HelpDeskService) ValidateHelpDesk(helpDeskID int) error {
	var exists bool
	if err := s.db.Model(&models.HelpDesk{}).Select("count(1) > 0").Where("id = ?", helpDeskID).Scan(&exists).Error; err != nil {
		return err
	}
	if !exists {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *HelpDeskService) ValidateTransaction(transactionID int) error {
	var exists bool
	if err := s.db.Model(&models.HelpDeskTransaction{}).Select("count(1) > 0").Where("id = ?", transactionID).Scan(&exists).Error; err != nil {
		return err
	}
	if !exists {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (s *HelpDeskService) ValidateUser(userID int) error {
	var exists bool
	if err := s.db.Model(&models.User{}).Select("count(1) > 0").Where("id = ?", userID).Scan(&exists).Error; err != nil {
		return err
	}
	if !exists {
		return gorm.ErrRecordNotFound
	}
	return nil
}

// EnsureUserAssignable verifies both help desk and user exist before linking.
func (s *HelpDeskService) EnsureUserAssignable(helpDeskID, userID int) error {
	if err := s.ValidateHelpDesk(helpDeskID); err != nil {
		return err
	}
	if err := s.ValidateUser(userID); err != nil {
		return err
	}
	return nil
}

// EnsureTransactionAssignable ensures transaction and user exist and transaction belongs to a help desk.
func (s *HelpDeskService) EnsureTransactionAssignable(transactionID, userID int) error {
	if err := s.ValidateTransaction(transactionID); err != nil {
		return err
	}
	if err := s.ValidateUser(userID); err != nil {
		return err
	}
	return nil
}
