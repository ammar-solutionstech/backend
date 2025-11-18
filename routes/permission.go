package routes

import (
	"backend/config"
	"backend/controllers"
	"backend/middleware"
	"net/http"

	"github.com/go-chi/chi"
)

func RegisterPermissionRoutes(mux *chi.Mux, cfg *config.Config) {

	mux.Route("/api/permissions", func(r chi.Router) {

		// ✅ Get all permissions
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View All Permissions", "Basic", cfg),
		).Get("/", http.HandlerFunc(controllers.GetPermissions))

		// ✅ Get one permission by ID
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Permission", "Basic", cfg),
		).Get("/{id}", http.HandlerFunc(controllers.GetPermission))

		// ✅ Create new permission
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Add Permission", "Basic", cfg),
		).Post("/", http.HandlerFunc(controllers.CreatePermission))

		// ✅ Update existing permission
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Edit Permission", "Basic", cfg),
		).Put("/{id}", http.HandlerFunc(controllers.UpdatePermission))

		// ✅ Delete permission
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Delete Permission", "Basic", cfg),
		).Delete("/{id}", http.HandlerFunc(controllers.DeletePermission))
	})
}
