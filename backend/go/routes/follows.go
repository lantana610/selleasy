package routes

import (
	"encoding/json"
	"net/http"

	"selleasy/models"
)

type createFollowRequest struct {
	UserID   string `json:"user_id"`
	Category string `json:"category"`
	City     string `json:"city"`
}

func RegisterFollowRoutes(mux *http.ServeMux, store *models.Store) {
	mux.HandleFunc("POST /api/v1/follows", func(w http.ResponseWriter, r *http.Request) {
		handleCreateFollow(w, r, store)
	})
}

func handleCreateFollow(w http.ResponseWriter, r *http.Request, store *models.Store) {
	var req createFollowRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.UserID == "" {
		http.Error(w, "user_id is required", http.StatusBadRequest)
		return
	}

	follow := models.Follow{
		UserID:   req.UserID,
		Category: req.Category,
		City:     req.City,
	}

	created := store.CreateFollow(follow)

	json.NewEncoder(w).Encode(created)
}