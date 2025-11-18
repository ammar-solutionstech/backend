package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"

	"backend/services"
)

var helpDeskRelationService = services.NewHelpDeskService(cnf.DB)

// Team membership handlers.

func GetTeamMembers(w http.ResponseWriter, r *http.Request) {
	teamID, err := parseIDParam(r, "teamId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid team id")
		return
	}

	if err := helpDeskRelationService.ValidateTeam(teamID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "team not found")
		return
	}

	members, err := helpDeskRelationService.ListTeamMembers(teamID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list team members")
		return
	}
	writeJSON(w, http.StatusOK, members)
}

func AddTeamMember(w http.ResponseWriter, r *http.Request) {
	teamID, err := parseIDParam(r, "teamId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid team id")
		return
	}

	var payload struct {
		UserID int `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.UserID <= 0 {
		writeError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	if err := helpDeskRelationService.ValidateTeam(teamID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "team not found")
		return
	}
	if err := helpDeskRelationService.ValidateUser(payload.UserID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusBadRequest
		}
		writeError(w, status, "user not found")
		return
	}

	if err := helpDeskRelationService.AddTeamMember(teamID, payload.UserID); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to add team member")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{
		"help_desk_team_id": teamID,
		"user_id":           payload.UserID,
	})
}

func RemoveTeamMember(w http.ResponseWriter, r *http.Request) {
	teamID, err := parseIDParam(r, "teamId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid team id")
		return
	}
	userID, err := parseIDParam(r, "userId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := helpDeskRelationService.RemoveTeamMember(teamID, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "team member link not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to remove team member")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// Help desk participant handlers.

func GetHelpDeskParticipants(w http.ResponseWriter, r *http.Request) {
	helpDeskID, err := parseIDParam(r, "helpDeskId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid help desk id")
		return
	}

	if err := helpDeskRelationService.ValidateHelpDesk(helpDeskID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "help desk ticket not found")
		return
	}

	participants, err := helpDeskRelationService.ListParticipants(helpDeskID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list help desk participants")
		return
	}
	writeJSON(w, http.StatusOK, participants)
}

func AddHelpDeskParticipant(w http.ResponseWriter, r *http.Request) {
	helpDeskID, err := parseIDParam(r, "helpDeskId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid help desk id")
		return
	}

	var payload struct {
		UserID int `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.UserID <= 0 {
		writeError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	if err := helpDeskRelationService.EnsureUserAssignable(helpDeskID, payload.UserID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusBadRequest
		}
		writeError(w, status, "invalid help desk or user reference")
		return
	}

	if err := helpDeskRelationService.AddParticipant(helpDeskID, payload.UserID); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to add participant")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{
		"help_desk_id": helpDeskID,
		"user_id":      payload.UserID,
	})
}

func RemoveHelpDeskParticipant(w http.ResponseWriter, r *http.Request) {
	helpDeskID, err := parseIDParam(r, "helpDeskId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid help desk id")
		return
	}
	userID, err := parseIDParam(r, "userId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := helpDeskRelationService.RemoveParticipant(helpDeskID, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "participant link not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to remove participant")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}

// Transaction user assignment handlers.

func GetTransactionUsers(w http.ResponseWriter, r *http.Request) {
	transactionID, err := parseIDParam(r, "transactionId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid transaction id")
		return
	}

	if err := helpDeskRelationService.ValidateTransaction(transactionID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusNotFound
		}
		writeError(w, status, "transaction not found")
		return
	}

	users, err := helpDeskRelationService.ListTransactionUsers(transactionID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list transaction users")
		return
	}
	writeJSON(w, http.StatusOK, users)
}

func AddTransactionUser(w http.ResponseWriter, r *http.Request) {
	transactionID, err := parseIDParam(r, "transactionId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid transaction id")
		return
	}

	var payload struct {
		UserID int `json:"user_id"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil || payload.UserID <= 0 {
		writeError(w, http.StatusBadRequest, "user_id is required")
		return
	}

	if err := helpDeskRelationService.EnsureTransactionAssignable(transactionID, payload.UserID); err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = http.StatusBadRequest
		}
		writeError(w, status, "invalid transaction or user reference")
		return
	}

	if err := helpDeskRelationService.AddTransactionUser(transactionID, payload.UserID); err != nil {
		writeError(w, http.StatusInternalServerError, "failed to add transaction user")
		return
	}
	writeJSON(w, http.StatusCreated, map[string]int{
		"transaction_id": transactionID,
		"user_id":        payload.UserID,
	})
}

func RemoveTransactionUser(w http.ResponseWriter, r *http.Request) {
	transactionID, err := parseIDParam(r, "transactionId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid transaction id")
		return
	}
	userID, err := parseIDParam(r, "userId")
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	if err := helpDeskRelationService.RemoveTransactionUser(transactionID, userID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			writeError(w, http.StatusNotFound, "transaction user link not found")
			return
		}
		writeError(w, http.StatusInternalServerError, "failed to remove transaction user")
		return
	}
	writeJSON(w, http.StatusNoContent, nil)
}
