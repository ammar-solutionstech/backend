package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"backend/config"
	"backend/services"
	"backend/utils"
)

type AuthController struct {
	Cfg  *config.Config
	User *services.UserService
}

func NewAuthController(cfg *config.Config, userService *services.UserService) *AuthController {
	return &AuthController{Cfg: cfg, User: userService}
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request",
		})
		return
	}
	req.Email = strings.TrimSpace(req.Email)
	user, err := c.User.Authenticate(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	permissions, _ := c.User.GetAllUserPermissions(user.ID)
	roles, _ := c.User.GetUserRoles(user.ID)
	token, err := utils.GenerateToken(c.Cfg, user.ID, roles, permissions)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "token generation failed:" + err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
