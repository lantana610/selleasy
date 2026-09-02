package models

import (
	"sync"
	"encoding/hex"
	"crypto/rand"
	"errors"
)

type Store struct {
	mu       sync.Mutex
	Users    map[string]User
	Listings map[string]Listing
	Orders   map[string]Order
}

func NewStore() *Store {
	return &Store{
		Users:    make(map[string]User),
		Listings: make(map[string]Listing),
		Orders:   make(map[string]Order),  
	}
}
func (s *Store) CreateOrder(order Order) (Order, error){
	s.mu.Lock()
	defer s.mu.Unlock()

	listing, exits := s.Listings[order.ListingID]
	if !exits{
		return Order{}, errors.New("listing not found")

	}
	if listing.Status != StatusAvailable{
		return Order{}, errors.New("listing not available")
	}
	order.ID = newID()
	order.Status = OrderRequested
	s.Orders[order.ID] = order

	listing.Status = StatusPending
    s.Listings[order.ListingID] = listing

	return order, nil
 }

func (s *Store) CreateListing(listing Listing) Listing {
	s.mu.Lock()
	defer s.mu.Unlock()

	listing.ID = newID()
	listing.Status = StatusAvailable
	s.Listings[listing.ID] = listing
	return listing
}

func (s *Store) CreateUser(user User) User{
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = newID()
	s.Users[user.ID] = user
	return user
}
func (s *Store) GetUserByEmail(email string) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, user := range s.Users {
		if user.Email == email {
			return user, true
		}
	}
	return User{}, false
}
func newID() string{
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}
func (s *Store) GetAllListings() []Listing {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := []Listing{}
	for _, listing := range s.Listings {
		result = append(result, listing)
	}
	return result
}