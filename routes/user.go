package routes

import (
	"backend/config"
	"backend/controllers"
	"backend/middleware"
	"net/http"

	"github.com/go-chi/chi"
)

func RegisterUserRoutes(mux *chi.Mux, cfg *config.Config) {

	mux.Route("/api/users", func(r chi.Router) {

		// ✅ Get all Users
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View All Users", "Basic", cfg),
		).Get("/", http.HandlerFunc(controllers.GetUsers))

		// ✅ Get one User by ID
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("View User", "Basic", cfg),
		).Get("/{id}", http.HandlerFunc(controllers.GetUser))

		// ✅ Create new User
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Add User", "Basic", cfg),
		).Post("/", http.HandlerFunc(controllers.CreateUser))

		// ✅ Update existing User
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Edit User", "Basic", cfg),
		).Put("/{id}", http.HandlerFunc(controllers.UpdateUser))

		// ✅ Delete User
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Delete User", "Basic", cfg),
		).Delete("/{id}", http.HandlerFunc(controllers.DeleteUser))

		// ✅ Activate or deactivate user by ID {“active”:”yes”}or{“acive”:”no”}
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Suspend User", "Basic", cfg),
		).Put("/{id}/suspend", http.HandlerFunc(controllers.SuspendUser))

		////////////////////// For Permissions ///////////////////////////////
		// ✅ Update User's Permissions (Replace the old with new)
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Permissions of User", "Basic", cfg),
		).Put("/{id}/permissions", http.HandlerFunc(controllers.UpdateUserPermissions))

		// ✅ Add one Permission to User
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Permissions of User", "Basic", cfg),
		).Post("/{id}/permissions", http.HandlerFunc(controllers.AddPermissionToUser))

		// ✅ Delete one Permission From User
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Permissions of User", "Basic", cfg),
		).Delete("/{id}/permissions", http.HandlerFunc(controllers.DeletePermissionFromUser))

		////////////////////// For Roles ///////////////////////////////
		// ✅ Update User's Roles (Replace the old with new)
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Roles of User", "Basic", cfg),
		).Put("/{id}/roles", http.HandlerFunc(controllers.UpdateUserRoles))

		// ✅ Add one Role to User
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Roles of User", "Basic", cfg),
		).Post("/{id}/roles", http.HandlerFunc(controllers.AddRoleToUser))

		// ✅ Delete one Role From User
		r.With(
			middleware.JWTAuth(cfg),
			middleware.PermissionMiddleware("Update Roles of User", "Basic", cfg),
		).Delete("/{id}/roles", http.HandlerFunc(controllers.DeleteRoleFromUser))
	})
}
