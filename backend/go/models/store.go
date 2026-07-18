package models

import (
	"sync"
	"encoding/hex"
	"crypto/rand"
)

type Store struct {
	mu       sync.Mutex
	Users    map[string]User
	Listings map[string]Listing
}

func NewStore() *Store {
	return &Store{
		Users:    make(map[string]User),
		Listings: make(map[string]Listing),
	}
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