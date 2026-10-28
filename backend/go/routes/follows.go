package routes

import (
	"encoding/json"
	"net/http"

	"selleasy/models"
)

type createFollowRequest struct {
	Category string `json:"category"`
	City     string `json:"city"`
}

func RegisterFollowRoutes(mux *http.ServeMux, store *models.Store) {
	mux.HandleFunc("POST /api/v1/follows", func(w http.ResponseWriter, r *http.Request) {
		handleCreateFollow(w, r, store)
	})
}

func handleCreateFollow(w http.ResponseWriter, r *http.Request, store *models.Store) {
	user, ok := requireUser(w, r, store)
	if !ok {
		return
	}

	var req createFollowRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Category == "" && req.City == "" {
		http.Error(w, "choose a category, a city, or both", http.StatusBadRequest)
		return
	}

	follow := models.Follow{
		UserID:   user.ID,
		Category: req.Category,
		City:     req.City,
	}

	created := store.CreateFollow(follow)

	writeJSON(w, http.StatusCreated, created)
}