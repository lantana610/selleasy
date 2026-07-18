package models

type ListingStatus string

const (
	StatusAvailable ListingStatus = "available"
	StatusPending   ListingStatus = "pending"
	StatusSold      ListingStatus = "sold"
)

type Listing struct {
	ID          string
	SellerID    string
	Title       string
	Description string
	Category    string
	Price       float64
	Currency    string
	City        string
	State       string
	Country     string
	Status      ListingStatus
}
