package models

type Follow struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	Category string `json:"category"`
	City     string `json:"city"`
}