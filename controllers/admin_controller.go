package controllers

import (
	"encoding/json"
	"net/http"

	"backend/middleware"
)

// Example protected handler that returns some admin-only info.
func AdminOnlyHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, _ := r.Context().Value(middleware.ContextUserID).(string)
		// For demo: return current user ID and maybe a list of users
		users := []string{}
		// in a real app we would call DB/service to fetch users, here we just return usernames in memory
		// this example uses services package to show how to call it (GetByUsername exists). For simplicity we'll skip heavy logic.
		_ = users
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"message": "welcome admin",
			"user_id": userID,
		})
	}
}

func AdminOnly(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(middleware.ContextUserID)
	json.NewEncoder(w).Encode(map[string]any{
		"message": "Welcome, admin!",
		"user_id": userID,
	})
}
