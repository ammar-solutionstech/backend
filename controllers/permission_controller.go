package controllers

import (
	"net/http"

	//"github.com/gin-gonic/gin"

	"backend/config"
	"backend/models"
	"encoding/json"
	"strconv"
	"strings"
)

var cnf = config.Load()

func GetPermissions(w http.ResponseWriter, r *http.Request) {
	var permissions []models.Permission
	err := cnf.DB.Find(&permissions).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(permissions)
}

func GetPermission(w http.ResponseWriter, r *http.Request) {
	_, idStr, _ := strings.Cut(r.URL.Path, "/permissions/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var p models.Permission
	err = cnf.DB.Where("id = ?", id).First(&p).Error
	if err != nil {
		http.Error(w, "Permission not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func CreatePermission(w http.ResponseWriter, r *http.Request) {
	var p models.Permission
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := cnf.DB.Create(&p).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func UpdatePermission(w http.ResponseWriter, r *http.Request) {
	_, idStr, _ := strings.Cut(r.URL.Path, "/permissions/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var p models.Permission
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	err = cnf.DB.Model(&p).Where("id = ?", id).Updates(models.Permission{Name: p.Name, Module: p.Module}).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	/* rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Permission not found", http.StatusNotFound)
		return
	} */
	p.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func DeletePermission(w http.ResponseWriter, r *http.Request) {
	_, idStr, _ := strings.Cut(r.URL.Path, "/permissions/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = cnf.DB.Delete(&models.Permission{}, id).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
