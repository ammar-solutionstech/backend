package routes

import (
	"net/http"

	"backend/config"
	"backend/controllers"
	"backend/middleware"

	"github.com/go-chi/chi"
)

// RegisterHelpDeskRoutes registers help desk domain endpoints beyond types.
func RegisterHelpDeskRoutes(mux *chi.Mux, cfg *config.Config) {
	registerResourceRoutes(mux, cfg, "/api/help-desk", controllers.NewHelpDeskResource(cfg.DB), ResourcePermissions{
		Module: "Help Desk",
		List:   "View Help Desk Tickets",
		View:   "View Help Desk Ticket",
		Create: "Create Help Desk Ticket",
		Update: "Update Help Desk Ticket",
		Delete: "Delete Help Desk Ticket",
	})

	registerResourceRoutes(mux, cfg, "/api/help-desk/ratings", controllers.NewHelpDeskRatingResource(cfg.DB), ResourcePermissions{
		Module: "Help Desk",
		List:   "View Help Desk Ratings",
		View:   "View Help Desk Rating",
		Create: "Create Help Desk Rating",
		Update: "Update Help Desk Rating",
		Delete: "Delete Help Desk Rating",
	})

	registerResourceRoutes(mux, cfg, "/api/help-desk/transactions", controllers.NewHelpDeskTransactionResource(cfg.DB), ResourcePermissions{
		Module: "Help Desk",
		List:   "View Help Desk Transactions",
		View:   "View Help Desk Transaction",
		Create: "Create Help Desk Transaction",
		Update: "Update Help Desk Transaction",
		Delete: "Delete Help Desk Transaction",
	})

	registerResourceRoutes(mux, cfg, "/api/help-desk/teams", controllers.NewTeamResource(cfg.DB), ResourcePermissions{
		Module: "Help Desk",
		List:   "View Help Desk Teams",
		View:   "View Help Desk Team",
		Create: "Create Help Desk Team",
		Update: "Update Help Desk Team",
		Delete: "Delete Help Desk Team",
	})

	registerResourceRoutes(mux, cfg, "/api/help-desk/transaction-types", controllers.NewTransactionTypeResource(cfg.DB), ResourcePermissions{
		Module: "Help Desk",
		List:   "View Help Desk Transaction Types",
		View:   "View Help Desk Transaction Type",
		Create: "Create Help Desk Transaction Type",
		Update: "Update Help Desk Transaction Type",
		Delete: "Delete Help Desk Transaction Type",
	})

	mux.Route("/api/help-desk/teams/{teamId}/members", func(r chi.Router) {
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Team Members", "Help Desk", cfg),
		).Get("/", http.HandlerFunc(controllers.GetTeamMembers))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Add Team Member", "Help Desk", cfg),
		).Post("/", http.HandlerFunc(controllers.AddTeamMember))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Remove Team Member", "Help Desk", cfg),
		).Delete("/{userId}", http.HandlerFunc(controllers.RemoveTeamMember))
	})

	mux.Route("/api/help-desk/{helpDeskId}/participants", func(r chi.Router) {
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Help Desk Participants", "Help Desk", cfg),
		).Get("/", http.HandlerFunc(controllers.GetHelpDeskParticipants))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Add Help Desk Participant", "Help Desk", cfg),
		).Post("/", http.HandlerFunc(controllers.AddHelpDeskParticipant))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Remove Help Desk Participant", "Help Desk", cfg),
		).Delete("/{userId}", http.HandlerFunc(controllers.RemoveHelpDeskParticipant))
	})

	mux.Route("/api/help-desk/transactions/{transactionId}/users", func(r chi.Router) {
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View Transaction Users", "Help Desk", cfg),
		).Get("/", http.HandlerFunc(controllers.GetTransactionUsers))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Add Transaction User", "Help Desk", cfg),
		).Post("/", http.HandlerFunc(controllers.AddTransactionUser))

		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Remove Transaction User", "Help Desk", cfg),
		).Delete("/{userId}", http.HandlerFunc(controllers.RemoveTransactionUser))
	})
}
