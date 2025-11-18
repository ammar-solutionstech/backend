package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"

	"backend/services"
)

var navigationRelationService = services.NewNavigationService(cnf.DB)

func GetMenuRoles(w http.ResponseWriter, r *http.Request) {
	menuID, err := parseIDParam(r, "menuId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid menu id")
		return
	}

	if err := navigationRelationService.ValidateMenu(menuID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "menu not found")
		return
	}

	roles, err := navigationRelationService.ListMenuRoles(menuID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list menu roles")
		return
	}
	writeJSON(w, http.StatusOK, roles)
}

func AddMenuRole(w http.ResponseWriter, r *http.Request) {
	menuID, err := parseIDParam(r, "menuId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid menu id")
		return
	}

	var payload struct {
		RoleID int `json:"role_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.RoleID <= 0 {
		writeError(w, http.StatusBadRequest, "role_id is required")
		return
	}

	if err := navigationRelationService.ValidateMenu(menuID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "menu not found")
		return
	}
	if err := navigationRelationService.ValidateRole(payload.RoleID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusBadRequest
		}
		writeError(w, status, "role not found")
		return
	}

	if err := navigationRelationService.AddMenuRole(menuID, payload.RoleID); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to add role to menu")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{
		"menu_id": menuID,
		"role_id": payload.RoleID,
	})
}

func RemoveMenuRole(w http.ResponseWriter, r *http.Request) {
	menuID, err := parseIDParam(r, "menuId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid menu id")
		return
	}
	roleID, err := parseIDParam(r, "roleId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid role id")
		return
	}

	if err := navigationRelationService.RemoveMenuRole(menuID, roleID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "menu role link not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to remove role from menu")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}
