package routes

import (
	"net/http"

	"github.com/go-chi/chi"

	"backend/config"
	"backend/controllers"
	"backend/middleware"
)

// ResourcePermissions holds permission strings for CRUD operations.
type ResourcePermissions struct {
	Module string
	List   string
	View   string
	Create string
	Update string
	Delete string
}

func registerResourceRoutes(mux *chi.Mux, cfg *config.Config, basePath string, ctrl controllers.RESTController, perms ResourcePermissions) {
	mux.Route(basePath, func(r chi.Router) {
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware(perms.List, perms.Module, cfg),
		).Get("/", http.HandlerFunc(ctrl.List))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware(perms.View, perms.Module, cfg),
		).Get("/{id}", http.HandlerFunc(ctrl.Get))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware(perms.Create, perms.Module, cfg),
		).Post("/", http.HandlerFunc(ctrl.Create))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware(perms.Update, perms.Module, cfg),
		).Put("/{id}", http.HandlerFunc(ctrl.Update))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware(perms.Delete, perms.Module, cfg),
		).Delete("/{id}", http.HandlerFunc(ctrl.Delete))
	})
}
