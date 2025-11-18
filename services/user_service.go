package services

import (
	"database/sql"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"backend/models"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	u := &models.User{}
	err := s.DB.Preload("Roles").Preload("Permissions").Where("work_email = ?", email).First(&u).Error
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invalid email or password") //ErrInvalidCredentials
		}
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, ErrInvalidCredentials
	}

	/* u.Roles, _ = s.GetUserRoles(u.ID)
	u.Permissions, _ = s.GetAllUserPermissions(u.ID) */
	return u, nil
}

func (s *UserService) GetUserRoles(userID int) ([]string, error) {
	query := `
		SELECT r.name
		FROM public."user_role" ur
		JOIN public."Role" r ON ur.role_id = r.id
		WHERE ur.user_id = ?`
	var roles []string
	err := s.DB.Raw(query, userID).Scan(&roles).Error
	if err != nil {
		return nil, err
	}
	return roles, nil
}

/* func (s *UserService) GetUserPermissions(userID int) ([]string, error) {
	rows, err := s.DB.Query(`
		SELECT r.name | '---' | r.module
		FROM public."user_special_permission" ur
		JOIN public."User" r ON ur.user_id = r.id
		WHERE ur.user_id = $1`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var permissions []string
	for rows.Next() {
		var r string
		rows.Scan(&r)
		permissions = append(permissions, strings.Trim(r, " "))
	}
	return permissions, nil
} */

func (s *UserService) GetAllUserPermissions(userID int) ([]int, error) {
	query := `
		SELECT usp.permission_id
		FROM public."user_special_permission" usp
		WHERE usp.user_id = ?
		union
		SELECT rp.permission_id
		FROM role_permission rp
		JOIN user_role ur ON ur.role_id=rp.role_id
		WHERE ur.user_id = ?`
	var permissions []int
	err := s.DB.Raw(query, userID, userID).Scan(&permissions).Error
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
