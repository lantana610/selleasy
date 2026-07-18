package routes

import (
	"encoding/json"
	"net/http"

	"selleasy/models"
)

type createListingRequest struct {
	SellerID    string  `json:"seller_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	Country     string  `json:"country"`
}

func RegisterListingRoutes(mux *http.ServeMux, store *models.Store) {
	mux.HandleFunc("POST /api/v1/listings", func(w http.ResponseWriter, r *http.Request) {
		handleCreateListing(w, r, store)
	})
}

func handleCreateListing(w http.ResponseWriter, r *http.Request, store *models.Store) {
	var req createListingRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Title == "" || req.SellerID == "" || req.Price <= 0 {
		http.Error(w, "seller_id, title, and a positive price are required", http.StatusBadRequest)
		return
	}

	listing := models.Listing{
		SellerID:    req.SellerID,
		Title:       req.Title,
		Description: req.Description,
		Category:    req.Category,
		Price:       req.Price,
		Currency:    req.Currency,
		City:        req.City,
		State:       req.State,
		Country:     req.Country,
	}

	created := store.CreateListing(listing)

	json.NewEncoder(w).Encode(created)
}
