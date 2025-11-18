package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"gorm.io/gorm"

	"backend/services"
)

// RESTController defines the standard CRUD handler set.
type RESTController interface {
	List(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
}

// ResourceController wires GenericService to HTTP handlers.
type ResourceController[T any] struct {
	service      *services.GenericService[T]
	resourceName string
	allowed      map[string]struct{}
}

// NewResourceController constructs a resource controller.
func NewResourceController[T any](db *gorm.DB, resource string, allowedFields []string) *ResourceController[T] {
	fieldSet := make(map[string]struct{}, len(allowedFields))
	for _, f := range allowedFields {
		fieldSet[f] = struct{}{}
	}
	return &ResourceController[T]{
		service:      services.NewGenericService[T](db),
		resourceName: resource,
		allowed:      fieldSet,
	}
}

// List handles GET collection.
func (c *ResourceController[T]) List(w http.ResponseWriter, r *http.Request) {
	var records []T
	if err := c.service.List(&records); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list "+c.resourceName+"s")
		return
	}
	writeJSON(w, http.StatusOK, records)
}

// Get handles GET by ID.
func (c *ResourceController[T]) Get(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var record T
	if err := c.service.Get(id, &record); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, c.resourceName+" not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to retrieve "+c.resourceName)
		return
	}
	writeJSON(w, http.StatusOK, record)
}

// Create handles POST.
func (c *ResourceController[T]) Create(w http.ResponseWriter, r *http.Request) {
	var payload T
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body : "+err.Error())
		return
	}

	if err := c.service.Create(&payload); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create "+c.resourceName)
		return
	}
	writeJSON(w, http.StatusCreated, payload)
}

// Update handles PUT.
func (c *ResourceController[T]) Update(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	filtered := c.filterAllowed(updates)
	if len(filtered) == 0 {
		writeError(w, http.StatusBadRequest, "no fields to update")
		return
	}

	record, err := c.service.Update(id, filtered)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, c.resourceName+" not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to update "+c.resourceName)
		return
	}
	writeJSON(w, http.StatusOK, record)
}

// Delete handles DELETE.
func (c *ResourceController[T]) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := parseIDParam(r, "id")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id parameter")
		return
	}
	if err := c.service.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, c.resourceName+" not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete "+c.resourceName)
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

func (c *ResourceController[T]) filterAllowed(updates map[string]interface{}) map[string]interface{} {
	if len(c.allowed) == 0 {
		return updates
	}
	filtered := make(map[string]interface{}, len(updates))
	for k, v := range updates {
		field := strings.TrimSpace(k)
		if field == "" || strings.EqualFold(field, "id") {
			continue
		}
		if _, ok := c.allowed[field]; ok {
			filtered[field] = v
		}
	}
	return filtered
}
