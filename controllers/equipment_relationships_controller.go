package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"gorm.io/gorm"

	"backend/models"
	"backend/services"
)

var equipmentRelationService = services.NewEquipmentService(cnf.DB)

// Equipment Software Handlers

func GetEquipmentSoftware(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	if err := equipmentRelationService.ValidateEquipmentExists(equipmentID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "equipment not found")
		return
	}

	links, err := equipmentRelationService.ListSoftware(equipmentID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list software")
		return
	}
	writeJSON(w, http.StatusOK, links)
}

func CreateEquipmentSoftware(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	var payload models.EquipmentSoftware
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	payload.EquipmentID = equipmentID

	if err := equipmentRelationService.ValidateEquipmentExists(payload.EquipmentID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "equipment not found")
		return
	}
	if err := equipmentRelationService.ValidateSoftwareExists(payload.SoftwareID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusBadRequest
		}
		writeError(w, status, "software not found")
		return
	}

	if err := equipmentRelationService.CreateSoftwareLink(&payload); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to link software to equipment")
		return
	}
	writeJSON(w, http.StatusCreated, payload)
}

func UpdateEquipmentSoftware(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	softwareID, err := parseIDParam(r, "softwareId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid software id")
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	delete(updates, "equipment_id")
	delete(updates, "software_id")

	record, err := equipmentRelationService.UpdateSoftwareLink(equipmentID, softwareID, updates)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "software assignment not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to update software assignment")
		return
	}
	writeJSON(w, http.StatusOK, record)
}

func DeleteEquipmentSoftware(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	softwareID, err := parseIDParam(r, "softwareId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid software id")
		return
	}

	if err := equipmentRelationService.DeleteSoftwareLink(equipmentID, softwareID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "software assignment not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete software assignment")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// Equipment Help Desk Links

func GetEquipmentHelpDeskLinks(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	if err := equipmentRelationService.ValidateEquipmentExists(equipmentID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "equipment not found")
		return
	}

	links, err := equipmentRelationService.ListHelpDeskLinks(equipmentID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list equipment help desk links")
		return
	}
	writeJSON(w, http.StatusOK, links)
}

func CreateEquipmentHelpDeskLink(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	var payload struct {
		HelpDeskID int `json:"help_desk_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.HelpDeskID <= 0 {
		writeError(w, http.StatusBadRequest, "help_desk_id is required")
		return
	}

	if err := equipmentRelationService.ValidateEquipmentExists(equipmentID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "equipment not found")
		return
	}
	if err := equipmentRelationService.ValidateHelpDeskExists(payload.HelpDeskID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusBadRequest
		}
		writeError(w, status, "help desk ticket not found")
		return
	}

	if err := equipmentRelationService.AddHelpDeskLink(equipmentID, payload.HelpDeskID); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to link equipment to help desk")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{
		"equipment_id": equipmentID,
		"help_desk_id": payload.HelpDeskID,
	})
}

func DeleteEquipmentHelpDeskLink(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}
	helpDeskID, err := parseIDParam(r, "helpDeskId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid help desk id")
		return
	}

	if err := equipmentRelationService.RemoveHelpDeskLink(equipmentID, helpDeskID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "link not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to unlink equipment from help desk")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// Equipment User History

func GetEquipmentUserHistory(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	if err := equipmentRelationService.ValidateEquipmentExists(equipmentID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "equipment not found")
		return
	}

	history, err := equipmentRelationService.ListUserHistory(equipmentID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list equipment user history")
		return
	}
	writeJSON(w, http.StatusOK, history)
}

func CreateEquipmentUserHistory(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	var payload models.EquipmentUserHistory
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	payload.EquipmentID = equipmentID

	if err := equipmentRelationService.ValidateEquipmentExists(payload.EquipmentID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "equipment not found")
		return
	}
	if err := equipmentRelationService.ValidateUserExists(payload.UserID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusBadRequest
		}
		writeError(w, status, "user not found")
		return
	}
	if err := services.EnsureHistoryWindow(payload.StartDate, payload.EndDate); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := equipmentRelationService.AddUserHistory(&payload); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create history entry")
		return
	}
	writeJSON(w, http.StatusCreated, payload)
}

func UpdateEquipmentUserHistory(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}

	userID, err := parseIDParam(r, "userId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	startDateRaw := chi.URLParam(r, "startDate")
	startDate, err := time.Parse("2006-01-02", startDateRaw)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid start date format, expected YYYY-MM-DD")
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	delete(updates, "equipment_id")
	delete(updates, "user_id")
	delete(updates, "start_date")

	if endRaw, ok := updates["end_date"].(string); ok {
		if parsed, err := time.Parse("2006-01-02", endRaw); err == nil {
			if err := services.EnsureHistoryWindow(startDate, parsed); err != nil {
				writeError(w, http.StatusBadRequest, err.Error())
				return
			}
		}
	}

	entry, err := equipmentRelationService.UpdateUserHistory(equipmentID, userID, startDate, updates)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "history entry not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to update history entry")
		return
	}
	writeJSON(w, http.StatusOK, entry)
}

func DeleteEquipmentUserHistory(w http.ResponseWriter, r *http.Request) {
	equipmentID, err := parseIDParam(r, "equipmentId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid equipment id")
		return
	}
	userID, err := parseIDParam(r, "userId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}
	startDateRaw := chi.URLParam(r, "startDate")
	startDate, err := time.Parse("2006-01-02", startDateRaw)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid start date format, expected YYYY-MM-DD")
		return
	}

	if err := equipmentRelationService.DeleteUserHistory(equipmentID, userID, startDate); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "history entry not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to delete history entry")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}
