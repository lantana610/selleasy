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

	log.Println("SellEasy API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
