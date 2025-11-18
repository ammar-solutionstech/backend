package routes

import (
	"net/http"

	"backend/config"
	"backend/controllers"
	"backend/middleware"

	"github.com/go-chi/chi"
)

// RegisterInventoryRoutes wires asset and inventory domain endpoints.
func RegisterInventoryRoutes(mux *chi.Mux, cfg *config.Config) {
	registerResourceRoutes(mux, cfg, "/api/brands", controllers.NewBrandResource(cfg.DB), ResourcePermissions{
		Module: "Inventory",
		List:   "View Brands",
		View:   "View Brand",
		Create: "Create Brand",
		Update: "Update Brand",
		Delete: "Delete Brand",
	})

	registerResourceRoutes(mux, cfg, "/api/models", controllers.NewModelResource(cfg.DB), ResourcePermissions{
		Module: "Inventory",
		List:   "View Models",
		View:   "View Model",
		Create: "Create Model",
		Update: "Update Model",
		Delete: "Delete Model",
	})

	registerResourceRoutes(mux, cfg, "/api/equipment-types", controllers.NewEquipmentTypeResource(cfg.DB), ResourcePermissions{
		Module: "Inventory",
		List:   "View Equipment Types",
		View:   "View Equipment Type",
		Create: "Create Equipment Type",
		Update: "Update Equipment Type",
		Delete: "Delete Equipment Type",
	})

	registerResourceRoutes(mux, cfg, "/api/operating-systems", controllers.NewOperatingSystemResource(cfg.DB), ResourcePermissions{
		Module: "Inventory",
		List:   "View Operating Systems",
		View:   "View Operating System",
		Create: "Create Operating System",
		Update: "Update Operating System",
		Delete: "Delete Operating System",
	})

	registerResourceRoutes(mux, cfg, "/api/software-categories", controllers.NewSoftwareCategoryResource(cfg.DB), ResourcePermissions{
		Module: "Inventory",
		List:   "View Software Categories",
		View:   "View Software Category",
		Create: "Create Software Category",
		Update: "Update Software Category",
		Delete: "Delete Software Category",
	})

	registerResourceRoutes(mux, cfg, "/api/software", controllers.NewSoftwareResource(cfg.DB), ResourcePermissions{
		Module: "Inventory",
		List:   "View Software",
		View:   "View Software",
		Create: "Create Software",
		Update: "Update Software",
		Delete: "Delete Software",
	})

	registerResourceRoutes(mux, cfg, "/api/equipment", controllers.NewEquipmentResource(cfg.DB), ResourcePermissions{
		Module: "Inventory",
		List:   "View Equipment",
		View:   "View Equipment",
		Create: "Create Equipment",
		Update: "Update Equipment",
		Delete: "Delete Equipment",
	})

	mux.Route("/api/equipment/{equipmentId}/software", func(r chi.Router) {
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Equipment Software", "Inventory", cfg),
		).Get("/", http.HandlerFunc(controllers.GetEquipmentSoftware))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Link Software To Equipment", "Inventory", cfg),
		).Post("/", http.HandlerFunc(controllers.CreateEquipmentSoftware))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Equipment Software", "Inventory", cfg),
		).Put("/{softwareId}", http.HandlerFunc(controllers.UpdateEquipmentSoftware))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Unlink Software From Equipment", "Inventory", cfg),
		).Delete("/{softwareId}", http.HandlerFunc(controllers.DeleteEquipmentSoftware))
	})

	mux.Route("/api/equipment/{equipmentId}/help-desk", func(r chi.Router) {
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Equipment Help Desk Links", "Inventory", cfg),
		).Get("/", http.HandlerFunc(controllers.GetEquipmentHelpDeskLinks))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Link Equipment To Help Desk", "Inventory", cfg),
		).Post("/", http.HandlerFunc(controllers.CreateEquipmentHelpDeskLink))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Unlink Equipment From Help Desk", "Inventory", cfg),
		).Delete("/{helpDeskId}", http.HandlerFunc(controllers.DeleteEquipmentHelpDeskLink))
	})

	mux.Route("/api/equipment/{equipmentId}/user-history", func(r chi.Router) {
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Equipment User History", "Inventory", cfg),
		).Get("/", http.HandlerFunc(controllers.GetEquipmentUserHistory))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Create Equipment User History", "Inventory", cfg),
		).Post("/", http.HandlerFunc(controllers.CreateEquipmentUserHistory))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Equipment User History", "Inventory", cfg),
		).Put("/{userId}/{startDate}", http.HandlerFunc(controllers.UpdateEquipmentUserHistory))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Delete Equipment User History", "Inventory", cfg),
		).Delete("/{userId}/{startDate}", http.HandlerFunc(controllers.DeleteEquipmentUserHistory))
	})

	registerResourceRoutes(mux, cfg, "/api/documents", controllers.NewDocumentResource(cfg.DB), ResourcePermissions{
		Module: "Inventory",
		List:   "View Documents",
		View:   "View Document",
		Create: "Create Document",
		Update: "Update Document",
		Delete: "Delete Document",
	})

	registerResourceRoutes(mux, cfg, "/api/maintenance", controllers.NewMaintenanceResource(cfg.DB), ResourcePermissions{
		Module: "Maintenance",
		List:   "View Maintenance Records",
		View:   "View Maintenance Record",
		Create: "Create Maintenance Record",
		Update: "Update Maintenance Record",
		Delete: "Delete Maintenance Record",
	})
}
