package routes

import (
	"backend/config"
	"backend/controllers"
	"backend/middleware"
	"net/http"

	"github.com/go-chi/chi"
)

func RegisterRoleRoutes(mux *chi.Mux, cfg *config.Config) {

	mux.Route("/api/roles", func(r chi.Router) {

		// ✅ Get all Roles
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View All Roles", "Basic", cfg),
		).Get("/", http.HandlerFunc(controllers.GetRoles))

		// ✅ Get one Role by ID
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Role", "Basic", cfg),
		).Get("/{id}", http.HandlerFunc(controllers.GetRole))

		// ✅ Create new Role
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Add Role", "Basic", cfg),
		).Post("/", http.HandlerFunc(controllers.CreateRole))

		// ✅ Update existing Role
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Edit Role", "Basic", cfg),
		).Put("/{id}", http.HandlerFunc(controllers.UpdateRole))

		// ✅ Delete Role
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Delete Role", "Basic", cfg),
		).Delete("/{id}", http.HandlerFunc(controllers.DeleteRole))

		// ✅ Update Role's Permissions (Replace the old with new)
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Permissions of Role", "Basic", cfg),
		).Put("/{id}/permissions", http.HandlerFunc(controllers.UpdateRolePermissions))

		// ✅ Add one Permission to Role
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Permissions of Role", "Basic", cfg),
		).Post("/{id}/permission", http.HandlerFunc(controllers.AddPermissionToRole))

		// ✅ Delete one Permission From Role
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Permissions of Role", "Basic", cfg),
		).Delete("/{id}/permission", http.HandlerFunc(controllers.DeletePermissionFromRole))
	})
}
