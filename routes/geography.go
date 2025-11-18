package routes

import (
	"backend/config"
	"backend/controllers"

	"github.com/go-chi/chi"
)

// RegisterGeographyRoutes wires geographic reference endpoints.
func RegisterGeographyRoutes(mux *chi.Mux, cfg *config.Config) {
	registerResourceRoutes(mux, cfg, "/api/countries", controllers.NewCountryResource(cfg.DB), ResourcePermissions{
		Module: "Directory",
		List:   "View Countries",
		View:   "View Country",
		Create: "Create Country",
		Update: "Update Country",
		Delete: "Delete Country",
	})

	registerResourceRoutes(mux, cfg, "/api/cities", controllers.NewCityResource(cfg.DB), ResourcePermissions{
		Module: "Directory",
		List:   "View Cities",
		View:   "View City",
		Create: "Create City",
		Update: "Update City",
		Delete: "Delete City",
	})

	registerResourceRoutes(mux, cfg, "/api/locations", controllers.NewLocationResource(cfg.DB), ResourcePermissions{
		Module: "Directory",
		List:   "View Locations",
		View:   "View Location",
		Create: "Create Location",
		Update: "Update Location",
		Delete: "Delete Location",
	})

	registerResourceRoutes(mux, cfg, "/api/contacts", controllers.NewContactResource(cfg.DB), ResourcePermissions{
		Module: "Directory",
		List:   "View Contacts",
		View:   "View Contact",
		Create: "Create Contact",
		Update: "Update Contact",
		Delete: "Delete Contact",
	})
}
