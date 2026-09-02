package models

type OrderStatus string

const (
	OrderRequested OrderStatus = "requested"
	OrderCompleted OrderStatus = "completed"
	OrderCancelled OrderStatus = "cancelled"
)

type Order struct {
	ID            string      `json:"id"`
	ListingID     string      `json:"listing_id"`
	BuyerID       string      `json:"buyer_id"`
	SellerID      string      `json:"seller_id"`
	Status        OrderStatus `json:"status"`
	PaymentMethod string      `json:"payment_method"`
	Amount        float64     `json:"amount"`
	Currency      string      `json:"currency"`
}