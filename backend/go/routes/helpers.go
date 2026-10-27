package routes

import (
	"encoding/json"
	"net/http"
	"strings"

	"selleasy/models"
)

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func requireUser(w http.ResponseWriter, r *http.Request, store *models.Store) (models.User, bool) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		http.Error(w, "please log in first", http.StatusUnauthorized)
		return models.User{}, false
	}
	user, ok := store.SessionUser(token)
	if !ok {
		http.Error(w, "your session has expired, please log in again", http.StatusUnauthorized)
		return models.User{}, false
	}
	return user, true
}