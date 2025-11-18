package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	"gorm.io/gorm"

	"backend/models"
	"backend/services"
)

var helpDeskTypeService = services.NewHelpDeskTypeService(cnf.DB)

type helpDeskTypeCreateRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}

type helpDeskTypeUpdateRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	IsActive    *bool   `json:"is_active"`
}

// GetHelpDeskTypes handles GET /api/help-desk-types
func GetHelpDeskTypes(w http.ResponseWriter, r *http.Request) {
	types, err := helpDeskTypeService.List()
	if err != nil {
		http.Error(w, "failed to retrieve help desk types", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types)
}

// GetHelpDeskType handles GET /api/help-desk-types/{id}
func GetHelpDeskType(w http.ResponseWriter, r *http.Request) {
	id, ok := parseHelpDeskTypeID(w, r)
	if !ok {
		return
	}

	helpDeskType, err := helpDeskTypeService.Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "help desk type not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to retrieve help desk type", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helpDeskType)
}

// CreateHelpDeskType handles POST /api/help-desk-types
func CreateHelpDeskType(w http.ResponseWriter, r *http.Request) {
	var payload helpDeskTypeCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	payload.Name = strings.TrimSpace(payload.Name)
	if payload.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	helpDeskType := models.HelpDeskType{
		Name: payload.Name,
	}

	if payload.Description != nil {
		helpDeskType.Description = strings.TrimSpace(*payload.Description)
	}
	helpDeskType.IsActive = payload.IsActive

	if err := helpDeskTypeService.Create(&helpDeskType); err != nil {
		http.Error(w, "failed to create help desk type", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(helpDeskType)
}

// UpdateHelpDeskType handles PUT /api/help-desk-types/{id}
func UpdateHelpDeskType(w http.ResponseWriter, r *http.Request) {
	id, ok := parseHelpDeskTypeID(w, r)
	if !ok {
		return
	}

	var payload helpDeskTypeUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	updates := make(map[string]interface{})

	if payload.Name != nil {
		name := strings.TrimSpace(*payload.Name)
		if name == "" {
			http.Error(w, "name cannot be empty", http.StatusBadRequest)
			return
		}
		updates["name"] = name
	}

	if payload.Description != nil {
		desc := strings.TrimSpace(*payload.Description)
		updates["description"] = desc
	}

	if payload.IsActive != nil {
		updates["is_active"] = *payload.IsActive
	}

	if len(updates) == 0 {
		http.Error(w, "no fields to update", http.StatusBadRequest)
		return
	}

	helpDeskType, err := helpDeskTypeService.Update(id, updates)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "help desk type not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to update help desk type", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(helpDeskType)
}

// DeleteHelpDeskType handles DELETE /api/help-desk-types/{id}
func DeleteHelpDeskType(w http.ResponseWriter, r *http.Request) {
	id, ok := parseHelpDeskTypeID(w, r)
	if !ok {
		return
	}

	if err := helpDeskTypeService.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "help desk type not found", http.StatusNotFound)
			return
		}
		http.Error(w, "failed to delete help desk type", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func parseHelpDeskTypeID(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return 0, false
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "invalid id parameter", http.StatusBadRequest)
		return 0, false
	}

	return id, true
}
