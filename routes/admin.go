package routes

import (
	"backend/config"
	"backend/controllers"
	"backend/middleware"
	"net/http"

	"github.com/go-chi/chi"
)

/* func RegisterAdminRoutes(mux *chi.Mux, cfg *config.Config) {
	// Protected admin route
	adminHandler := http.HandlerFunc(controllers.AdminOnly)
	adminProtected := middleware.JWTAuth(cfg, middleware.RequireRoles("admin")(adminHandler))
	mux.Handle("/api/admin", adminProtected)
} */

func RegisterAdminRoutes(mux *chi.Mux, cfg *config.Config) {

	mux.Route("/admin", func(r chi.Router) {

		// ✅ Get all permissions
		r.With(
			middleware.JWTAuth(cfg),
			middleware.RequireRoles("admin"),
		).Get("/", http.HandlerFunc(controllers.AdminOnly))

	})
}
