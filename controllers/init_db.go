package controllers

import (
	"backend/config"
	"backend/services"
)

type InitController struct {
	Cfg  *config.Config
	User *services.UserService
}

/* func (c *AuthController) createUser(w http.ResponseWriter, r *http.Request) {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}
	req.Email = strings.TrimSpace(req.Email)

	user, err := c.User.Authenticate(req.Email, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized) //.append("invalid credentials 3")
		return
	}
	fmt.Println(user)
	token, err := utils.GenerateToken(c.Cfg, user.ID, user.Roles, user.Permissions)
	if err != nil {
		http.Error(w, "token generation failed", http.StatusInternalServerError)
		return
	}
	fmt.Println(token)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
} */
