package models

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"sync"
	"os"
	"encoding/json"
)

type Store struct {
	mu       sync.Mutex
	Users    map[string]User
	Listings map[string]Listing
	Orders   map[string]Order
	Follows   map[string]Follow
}

func NewStore() *Store {
	return &Store{
		Users:    make(map[string]User),
		Listings: make(map[string]Listing),
		Orders:   make(map[string]Order),
		Follows:   make(map[string]Follow),
	}
}

func (s *Store) GetMatchingFollowers(listing Listing) []Follow {
	s.mu.Lock()
	defer s.mu.Unlock()

	matches := []Follow{}
	for _, follow := range s.Follows {
		categoryMatches := follow.Category == "" || follow.Category == listing.Category
		cityMatches := follow.City == "" || follow.City == listing.City

		if categoryMatches && cityMatches {
			matches = append(matches, follow)
		}
	}
	return matches
}
func (s *Store) GetUser(id string) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.Users[id]
	return user, exists
}
func (s *Store) CreateOrder(order Order) (Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	listing, exits := s.Listings[order.ListingID]
	if !exits {
		return Order{}, errors.New("listing not found")

	}
	if listing.Status != StatusAvailable {
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
func (s *Store) GetListing(id string) (Listing, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	listing, exists := s.Listings[id]
	return listing, exists
}

func (s *Store) CreateUser(user User) User {
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
func (s *Store) UpdateOrderStatus(orderID string, status OrderStatus) (Order, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, exists := s.Orders[orderID]
	if !exists {
		return Order{}, errors.New("order not found")
	}

	order.Status = status

	s.Orders[orderID] = order
	listing, exists := s.Listings[order.ListingID]
	if exists {
		if status == OrderCompleted {
			listing.Status = StatusSold
		} else if status == OrderCancelled {
			listing.Status = StatusAvailable
		}
		s.Listings[order.ListingID] = listing
	}
	return order, nil
}
func newID() string {
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
func (s *Store) CreateFollow(follow Follow) Follow {
	s.mu.Lock()
	defer s.mu.Unlock()

	follow.ID = newID()
	s.Follows[follow.ID] = follow
	return follow
}

func (s *Store) SaveToFile(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func (s *Store) LoadFromFile(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	return json.Unmarshal(data, s)
}