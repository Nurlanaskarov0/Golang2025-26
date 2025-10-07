package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type GetUserResponse struct {
	UserID int `json:"user_id"`
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

type CreateUserResponse struct {
	Created string `json:"created"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid id"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid id"})
		return
	}

	writeJSON(w, http.StatusOK, GetUserResponse{UserID: id})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid name"})
		return
	}

	// Validate name
	if strings.TrimSpace(req.Name) == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid name"})
		return
	}

	writeJSON(w, http.StatusCreated, CreateUserResponse{Created: req.Name})
}
