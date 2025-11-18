package routes

import (
	"net/http"

	"backend/config"
	"backend/controllers"
	"backend/middleware"

	"github.com/go-chi/chi"
)

// RegisterNavigationRoutes wires menu administration endpoints.
func RegisterNavigationRoutes(mux *chi.Mux, cfg *config.Config) {
	registerResourceRoutes(mux, cfg, "/api/menus", controllers.NewMenuResource(cfg.DB), ResourcePermissions{
		Module: "Navigation",
		List:   "View Menus",
		View:   "View Menu",
		Create: "Create Menu",
		Update: "Update Menu",
		Delete: "Delete Menu",
	})

	mux.Route("/api/menus/{menuId}/roles", func(r chi.Router) {
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Menu Roles", "Navigation", cfg),
		).Get("/", http.HandlerFunc(controllers.GetMenuRoles))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Add Menu Role", "Navigation", cfg),
		).Post("/", http.HandlerFunc(controllers.AddMenuRole))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Remove Menu Role", "Navigation", cfg),
		).Delete("/{roleId}", http.HandlerFunc(controllers.RemoveMenuRole))
	})
}
