package routes

import (
	"encoding/json"
	"net/http"

	"selleasy/models"
)

type placeOrderRequest struct {
	ListingID     string  `json:"listing_id"`
	BuyerID       string  `json:"buyer_id"`
	PaymentMethod string  `json:"payment_method"`
	Amount        float64 `json:"amount"`
	Currency      string  `json:"currency"`
}

type updateOrderStatusRequest struct {
	Status string `json:"status"`
}

func RegisterOrderRoutes(mux *http.ServeMux, store *models.Store) {
	mux.HandleFunc("POST /api/v1/orders", func(w http.ResponseWriter, r *http.Request) {
		handlePlaceOrder(w, r, store)
	})
	mux.HandleFunc("PATCH /api/v1/orders/{id}/status", func(w http.ResponseWriter, r *http.Request) {
		handleUpdateOrderStatus(w, r, store)
	})
}

func handlePlaceOrder(w http.ResponseWriter, r *http.Request, store *models.Store) {
	var req placeOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	if req.ListingID == "" || req.BuyerID == "" {
		http.Error(w, "listing_id and buyer_id are required", http.StatusBadRequest)
		return
	}

	listing, exists := store.GetListing(req.ListingID)
	if !exists {
		http.Error(w, "listing not found", http.StatusNotFound)
		return
	}

	order := models.Order{
		ListingID:     req.ListingID,
		BuyerID:       req.BuyerID,
		SellerID:      listing.SellerID,
		PaymentMethod: req.PaymentMethod,
		Amount:        req.Amount,
		Currency:      req.Currency,
	}

	created, err := store.CreateOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	json.NewEncoder(w).Encode(created)
}

func handleUpdateOrderStatus(w http.ResponseWriter, r *http.Request, store *models.Store) {
	orderID := r.PathValue("id")

	var req updateOrderStatusRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	status := models.OrderStatus(req.Status)
	if status != models.OrderCompleted && status != models.OrderCancelled && status != models.OrderRequested {
		http.Error(w, "status must be requested, completed, or cancelled", http.StatusBadRequest)
		return
	}

	updated, err := store.UpdateOrderStatus(orderID, status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(updated)
}