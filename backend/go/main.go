package main

import (
	"log"
	"net/http"
	"selleasy/models"
	"selleasy/routes"
)


func main() {

	store := models.NewStore()

	mux := http.NewServeMux()

	pattern := "GET /health"

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}

	mux.HandleFunc(pattern, handler)
	routes.RegisterAuthRoutes(mux, store)
	routes.RegisterListingRoutes(mux, store)
	routes.RegisterOrderRoutes(mux, store)

	log.Println("SellEasy API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(mux)))
}
func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
