package controllers

import (
	"fmt"
	"net/http"

	//"github.com/gin-gonic/gin"

	"backend/models"
	"encoding/json"
	"strconv"
	"strings"
)

// Retrive All Users
// `GET: /api/users`
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var Users []models.User
	err := cnf.DB.Preload("Nationality").Find(&Users).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't retrive Users : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Users)
}

// Retrive specific user
// `GET: /api/users/{user_id}`
func GetUser(w http.ResponseWriter, r *http.Request) {
	_, idStr, _ := strings.Cut(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}

	var user models.User
	err = cnf.DB.Preload("Teams").Preload("Nationality").Preload("Roles").Preload("Permissions").Where("id = ?", id).First(&user).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found : " + err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Create new user
// `POST: /api/users`
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input : " + err.Error(),
		})
		return
	}
	err := cnf.DB.Create(&user).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't create user : " + err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Update existing user
// `PUT: /api/users/{user_id}`
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	_, idStr, _ := strings.Cut(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input : " + err.Error(),
		})
		return
	}
	err = cnf.DB.Model(&user).Where("id = ?", id).Updates(user).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Update user : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Activate or deactivate user by ID
// `PUT: /api/users/{user_id}/suspend`
// body : {"active":"yes"}or{"acive":"no"}
func SuspendUser(w http.ResponseWriter, r *http.Request) {
	_, idWithPermission, _ := strings.Cut(r.URL.Path, "/users/")
	idStr, _, _ := strings.Cut(idWithPermission, "/suspend")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}
	var active struct {
		Active string `json:"active"`
	}
	var user models.User
	cnf.DB.Where("id = ?", id).First(&user)
	if err := json.NewDecoder(r.Body).Decode(&active); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input data : the input must be json as {\"active\":\"yes\"} or {\"active\":\"no\"}",
		})
		return
	}
	fmt.Println(active)
	if active.Active == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input data : the input must be json as {\"active\":\"yes\"} or {\"active\":\"no\"}",
		})
		return
	}
	active_value := 2
	if strings.ToLower(active.Active) == "yes" {
		active_value = 1
	}
	err = cnf.DB.Model(&user).Where("id = ?", id).Updates(models.User{Active: &active_value}).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Update user : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Delete specific user
// `DELETE: /api/users/{user_id}`
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	_, idStr, _ := strings.Cut(r.URL.Path, "/users/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}
	var user models.User
	err = cnf.DB.Preload("Nationality").Preload("Roles").Preload("Permissions").Where("id = ?", id).First(&user).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found : " + err.Error(),
		})
		return
	}

	err = cnf.DB.Delete(&models.User{}, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Delete User : " + err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User was deleted successfully",
		"user":    user,
	})
	w.WriteHeader(http.StatusOK)
}

// Update All user's additional permissions
// `PUT: /api/users/{user_id}/permissions`
// Input: []int // Array of permission's ids
func UpdateUserPermissions(w http.ResponseWriter, r *http.Request) {
	_, idWithPermission, _ := strings.Cut(r.URL.Path, "/users/")
	idStr, _, _ := strings.Cut(idWithPermission, "/permissions")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}

	var permIDs []int
	if err := json.NewDecoder(r.Body).Decode(&permIDs); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input : " + err.Error(),
		})
		return
	}
	var permissions []models.Permission
	if err := cnf.DB.Where("id IN ?", permIDs).Select("id, name, module").Find(&permissions).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid permissions : " + err.Error(),
		})
		return
	}
	var user_permissions []models.UserPermission
	for _, permission := range permissions {
		user_permissions = append(user_permissions, models.UserPermission{PermissionID: permission.ID, UserID: id})
	}
	err = cnf.DB.Where("user_id = ?", id).Delete(&user_permissions).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't delete User's permissions : " + err.Error(),
		})
		return
	}
	err = cnf.DB.Create(&user_permissions).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Add User's permissions : " + err.Error(),
		})
		return
	}
	var user models.User
	err = cnf.DB.Preload("Permissions").First(&user, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Permissions was updated successfully",
		"user":    user,
	})
}

// Add new permission to user
// `POST: /api/users/{user_id}/permissions`
// Input: {"id":int} // json of permission's id as "{"id":1}"
func AddPermissionToUser(w http.ResponseWriter, r *http.Request) {
	_, idWithPermission, _ := strings.Cut(r.URL.Path, "/users/")
	idStr, _, _ := strings.Cut(idWithPermission, "/permission")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}

	var permission models.Permission
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input : " + err.Error(),
		})
		return
	}
	if err := cnf.DB.Select("id, name, module").First(&permission, permission.ID).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid permissions : " + err.Error(),
		})
		return
	}
	fmt.Println(permission)
	user_permissions := models.UserPermission{PermissionID: permission.ID, UserID: id}
	fmt.Println(user_permissions)
	err = cnf.DB.Where(user_permissions).Assign(user_permissions).FirstOrCreate(&user_permissions).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Create User's Permission : " + err.Error(),
		})
		return
	}
	fmt.Println(user_permissions)
	var user models.User
	err = cnf.DB.Preload("Permissions").First(&user, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Permission was added successfully to user",
		"user":    user,
	})
}

// Delete one permission from user
// `DELETE: /api/users/{user_id}/permission`
// Input: {"id":int} // json of permission's id as "{"id":1}"
func DeletePermissionFromUser(w http.ResponseWriter, r *http.Request) {
	_, idWithPermission, _ := strings.Cut(r.URL.Path, "/users/")
	idStr, _, _ := strings.Cut(idWithPermission, "/permission")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}

	var permission models.Permission
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input : " + err.Error(),
		})
		return
	}
	user_permissions := models.UserPermission{PermissionID: permission.ID, UserID: id}
	err = cnf.DB.Where("user_id = ?", id).Where("permission_id = ?", permission.ID).Delete(&user_permissions).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Delete User's Permission : " + err.Error(),
		})
		return
	}
	var user models.User
	err = cnf.DB.Preload("Permissions").First(&user, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Permission was deleted successfully from user",
		"user":    user,
	})
}

// Update All user's roles
// `PUT: /api/users/{user_id}/roles`
// Input: []int // Array of role's ids
func UpdateUserRoles(w http.ResponseWriter, r *http.Request) {
	_, idWithRole, _ := strings.Cut(r.URL.Path, "/users/")
	idStr, _, _ := strings.Cut(idWithRole, "/roles")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}

	var roleIDs []int
	if err := json.NewDecoder(r.Body).Decode(&roleIDs); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input : " + err.Error(),
		})
		return
	}
	var roles []models.Role
	if err := cnf.DB.Where("id IN ?", roleIDs).Select("id, name, description").Find(&roles).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid roles : " + err.Error(),
		})
		return
	}
	if len(roles) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid roles provided",
		})
		return
	}
	var user_roles []models.UserRole
	for _, role := range roles {
		user_roles = append(user_roles, models.UserRole{RoleID: role.ID, UserID: id})
	}
	err = cnf.DB.Where("user_id = ?", id).Delete(&user_roles).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't delete User's roles : " + err.Error(),
		})
		return
	}
	err = cnf.DB.Create(&user_roles).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Add User's roles : " + err.Error(),
		})
		return
	}
	var user models.User
	err = cnf.DB.Preload("Roles").First(&user, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Roles was updated successfully",
		"user":    user,
	})
}

// Add new role to user
// `POST: /api/users/{user_id}/roles`
// Input: {"id":int} // json of role's id as "{"id":1}"
func AddRoleToUser(w http.ResponseWriter, r *http.Request) {
	_, idWithRole, _ := strings.Cut(r.URL.Path, "/users/")
	idStr, _, _ := strings.Cut(idWithRole, "/role")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}

	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input : " + err.Error(),
		})
		return
	}
	if err := cnf.DB.Select("id, name, description").First(&role, role.ID).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid roles : " + err.Error(),
		})
		return
	}

	user_roles := models.UserRole{RoleID: role.ID, UserID: id}
	err = cnf.DB.Where(user_roles).Assign(user_roles).FirstOrCreate(&user_roles).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Create User's Role : " + err.Error(),
		})
		return
	}
	var user models.User
	err = cnf.DB.Preload("Roles").First(&user, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Role was added successfully to user",
		"user":    user,
	})
}

// Delete one role from user
// `DELETE: /api/users/{user_id}/roles`
// Input: {"id":int} // json of role's id as "{"id":1}"
func DeleteRoleFromUser(w http.ResponseWriter, r *http.Request) {
	_, idWithRole, _ := strings.Cut(r.URL.Path, "/users/")
	idStr, _, _ := strings.Cut(idWithRole, "/role")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}

	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input : " + err.Error(),
		})
		return
	}
	if err := cnf.DB.Select("id, name, description").First(&role, role.ID).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid roles : " + err.Error(),
		})
		return
	}

	user_roles := models.UserRole{RoleID: role.ID, UserID: id}
	err = cnf.DB.Where("user_id = ?", id).Where("role_id = ?", role.ID).Delete(&user_roles).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Delete User's Role : " + err.Error(),
		})
		return
	}
	var user models.User
	err = cnf.DB.Preload("Roles").First(&user, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "User not found : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Role was deleted successfully from user",
		"user":    user,
	})
}
