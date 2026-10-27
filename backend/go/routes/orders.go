package routes

import (
	"encoding/json"
	"net/http"

	"selleasy/models"
	"selleasy/services"
)

type placeOrderRequest struct {
	ListingID     string `json:"listing_id"`
	PaymentMethod string `json:"payment_method"`
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
	mux.HandleFunc("GET /api/v1/orders", func(w http.ResponseWriter, r *http.Request) {
		handleGetOrders(w, r, store)
	})
}

func handlePlaceOrder(w http.ResponseWriter, r *http.Request, store *models.Store) {
	
	user, ok := requireUser(w, r, store)
	if !ok {
		return
	}

	var req placeOrderRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	if req.ListingID == "" {
		http.Error(w, "listing_id is required", http.StatusBadRequest)
		return
	}

	listing, exists := store.GetListing(req.ListingID)
	if !exists {
		http.Error(w, "listing not found", http.StatusNotFound)
		return
	}
	if listing.SellerID == user.ID {
		http.Error(w, "you can't order your own listing", http.StatusBadRequest)
		return
	}

	order := models.Order{
		ListingID:     req.ListingID,
		SellerID:      listing.SellerID,
		PaymentMethod: req.PaymentMethod,
		BuyerID:       user.ID,
		Amount:        listing.Price,
		Currency:      listing.Currency,
	}

	created, err := store.CreateOrder(order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	seller, exists := store.GetUser(created.SellerID)
	sellerEmail := ""
	if exists {
		sellerEmail = seller.Email
	}

	services.QueueEmailNotification(created.SellerID, "new_order", map[string]interface{}{
		"title": listing.Title,
		"email": sellerEmail,
	})
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
func handleGetOrders(w http.ResponseWriter, r *http.Request, store *models.Store) {
	buyerID := r.URL.Query().Get("buyer_id")
	sellerID := r.URL.Query().Get("seller_id")

	if buyerID != "" {
		orders := store.GetOrdersByBuyer(buyerID)
		json.NewEncoder(w).Encode(orders)
		return
	}

	if sellerID != "" {
		orders := store.GetOrdersBySeller(sellerID)
		json.NewEncoder(w).Encode(orders)
		return
	}

	http.Error(w, "buyer_id or seller_id is required", http.StatusBadRequest)
}
