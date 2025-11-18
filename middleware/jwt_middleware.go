package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"backend/config"
	"backend/utils"
)

// key type for context to avoid collisions
type contextKey string

const (
	ContextUserID      contextKey = "userID"
	ContextRoles       contextKey = "roles"
	ContextPermissions contextKey = "permissions"
)

func JWTAuth(cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if !strings.HasPrefix(auth, "Bearer ") {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "missing or invalid authorization header",
				})
				return
			}
			tokenStr := strings.TrimPrefix(auth, "Bearer ")
			claims, err := utils.ParseAndValidateToken(cfg, tokenStr)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "invalid token",
				})
				return
			}
			ctx := context.WithValue(r.Context(), ContextUserID, claims.UserID)
			ctx = context.WithValue(ctx, ContextRoles, claims.Roles)
			ctx = context.WithValue(ctx, ContextPermissions, claims.Permissions)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
