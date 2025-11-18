package routes

import (
	"net/http"

	"github.com/go-chi/chi"

	"backend/config"
	"backend/controllers"
	"backend/middleware"
)

// RegisterHelpDeskTypeRoutes wires help desk type endpoints with authentication and authorization guards.
func RegisterHelpDeskTypeRoutes(mux *chi.Mux, cfg *config.Config) {
	mux.Route("/api/help-desk-types", func(r chi.Router) {
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Help Desk Types", "Help Desk", cfg),
		).Get("/", http.HandlerFunc(controllers.GetHelpDeskTypes))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Help Desk Types", "Help Desk", cfg),
		).Get("/{id}", http.HandlerFunc(controllers.GetHelpDeskType))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Create Help Desk Type", "Help Desk", cfg),
		).Post("/", http.HandlerFunc(controllers.CreateHelpDeskType))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Help Desk Type", "Help Desk", cfg),
		).Put("/{id}", http.HandlerFunc(controllers.UpdateHelpDeskType))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Delete Help Desk Type", "Help Desk", cfg),
		).Delete("/{id}", http.HandlerFunc(controllers.DeleteHelpDeskType))
	})
}
