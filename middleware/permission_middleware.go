package middleware

import (
	"backend/config"
	"backend/models"
	"encoding/json"
	"net/http"
	"slices"
)

func PermissionMiddleware(requiredPermission string, module string, cfg *config.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			p := &models.Permission{}
			err := cfg.DB.Where("name = ? AND module = ?", requiredPermission, module).First(&p).Error
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Invalid Permission - " + requiredPermission + ":" + module,
				})
				return
			}

			roles, ok := r.Context().Value(ContextRoles).([]string)
			if ok {
				if slices.Contains(roles, "admin") {
					next.ServeHTTP(w, r)
					return
				}
			}

			perms, ok := r.Context().Value(ContextPermissions).([]int)
			if !ok {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "Permissions missing",
				})
				return
			}

			if slices.Contains(perms, p.ID) {
				next.ServeHTTP(w, r)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "Forbidden: missing permission " + requiredPermission,
			})
		})
	}
}
