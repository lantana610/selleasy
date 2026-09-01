package models

type ListingStatus string

const (
	StatusAvailable ListingStatus = "available"
	StatusPending   ListingStatus = "pending"
	StatusSold      ListingStatus = "sold"
)

type Listing struct {
	ID          string        `json:"id"`
	SellerID    string        `json:"seller_id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	Price       float64       `json:"price"`
	Currency    string        `json:"currency"`
	City        string        `json:"city"`
	State       string        `json:"state"`
	Country     string        `json:"country"`
	Status      ListingStatus `json:"status"`
}