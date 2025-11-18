package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"backend/config"
)

type CustomClaims struct {
	UserID      int      `json:"sub"`
	Roles       []string `json:"roles"`
	Permissions []int    `json:"permissions"`
	jwt.RegisteredClaims
}

func GenerateToken(cfg *config.Config, userID int, roles []string, permissions []int) (string, error) {
	now := time.Now().UTC()
	claims := CustomClaims{
		UserID:      userID,
		Roles:       roles,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.JWTIssuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(cfg.JWTExpiry)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWTSecret))
}

func ParseAndValidateToken(cfg *config.Config, tokenStr string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
