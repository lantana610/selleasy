package routes

import (
	"net/http"

	"selleasy/models"
)

func RegisterAuthRoutes(mux *http.ServeMux, store *models.Store) {
	mux.HandleFunc("POST /api/v1/register", func(w http.ResponseWriter, r *http.Request) {
		handleRegister(w, r, store)
	})
}

func handleRegister(w http.ResponseWriter, r *http.Request, store *models.Store) {
	w.Write([]byte("register endpoint hit"))
}