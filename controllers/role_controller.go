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

func GetRoles(w http.ResponseWriter, r *http.Request) {
	var Roles []models.Role
	err := cnf.DB.Find(&Roles).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't retrive Roles : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Roles)
}

func GetRole(w http.ResponseWriter, r *http.Request) {
	_, idStr, _ := strings.Cut(r.URL.Path, "/roles/")
	id, err := strconv.Atoi(idStr)
	fmt.Println(id)
	fmt.Println(err)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}

	var role models.Role
	err = cnf.DB.Preload("Permissions").Where("id = ?", id).First(&role).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Role not found : " + err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

func CreateRole(w http.ResponseWriter, r *http.Request) {
	var role models.Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid input : " + err.Error(),
		})
		return
	}
	err := cnf.DB.Create(&role).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't create Role : " + err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(role)
}

func UpdateRole(w http.ResponseWriter, r *http.Request) {
	_, idStr, _ := strings.Cut(r.URL.Path, "/roles/")
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
	err = cnf.DB.Model(&role).Where("id = ?", id).Updates(models.Role{Name: role.Name, Description: role.Description}).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Update Role : " + err.Error(),
		})
		return
	}
	role.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

func DeleteRole(w http.ResponseWriter, r *http.Request) {
	_, idStr, _ := strings.Cut(r.URL.Path, "/roles/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Invalid ID : " + err.Error(),
		})
		return
	}
	err = cnf.DB.Delete(&models.Role{}, id).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Delete Role : " + err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func UpdateRolePermissions(w http.ResponseWriter, r *http.Request) {
	_, idWithPermission, _ := strings.Cut(r.URL.Path, "/roles/")
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
	var role_permissions []models.RolePermission
	for _, permission := range permissions {
		role_permissions = append(role_permissions, models.RolePermission{PermissionID: permission.ID, RoleID: id})
	}
	err = cnf.DB.Where("role_id = ?", id).Delete(&role_permissions).Error //Where("permission_id IN ?", permIDs).
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't delete Role's permissions : " + err.Error(),
		})
		return
	}
	err = cnf.DB.Create(&role_permissions).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Add Role's permissions : " + err.Error(),
		})
		return
	}
	var role models.Role
	err = cnf.DB.Preload("Permissions").First(&role, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Role not found : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

func AddPermissionToRole(w http.ResponseWriter, r *http.Request) {
	_, idWithPermission, _ := strings.Cut(r.URL.Path, "/roles/")
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
	role_permissions := models.RolePermission{PermissionID: permission.ID, RoleID: id}
	err = cnf.DB.Create(&role_permissions).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Create Role's Permission : " + err.Error(),
		})
		return
	}
	var role models.Role
	err = cnf.DB.Preload("Permissions").First(&role, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Role not found : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}

func DeletePermissionFromRole(w http.ResponseWriter, r *http.Request) {
	_, idWithPermission, _ := strings.Cut(r.URL.Path, "/roles/")
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
	role_permissions := models.RolePermission{PermissionID: permission.ID, RoleID: id}
	err = cnf.DB.Where("role_id = ?", id).Where("permission_id = ?", permission.ID).Delete(&role_permissions).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Can't Delete Role's Permission : " + err.Error(),
		})
		return
	}
	var role models.Role
	err = cnf.DB.Preload("Permissions").First(&role, id).Error
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Role not found : " + err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(role)
}
