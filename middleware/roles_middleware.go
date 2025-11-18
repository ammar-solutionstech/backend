package middleware

import (
	"net/http"
	"strings"
)

func RequireRoles(required ...string) func(http.Handler) http.Handler {
	roleSet := make(map[string]struct{})
	for _, r := range required {
		roleSet[strings.ToLower(r)] = struct{}{}
	}
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			val := r.Context().Value(ContextRoles)
			roles, ok := val.([]string)
			if !ok {
				http.Error(w, "forbidden", http.StatusForbidden)
				return
			}
			for _, userRole := range roles {
				if _, ok := roleSet[strings.ToLower(userRole)]; ok {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "forbidden", http.StatusForbidden)
		})
	}
}
