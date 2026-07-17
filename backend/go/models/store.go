package models

import (
	"sync"
	"encoding/hex"
	"crypto/rand"
)

type Store struct {
	mu    sync.Mutex
	Users map[string]User
}
func NewStore() *Store {
	return &Store{
		Users: make(map[string]User),
	}
}

func (s *Store) CreateUser(user User) User{
	s.mu.Lock()
	defer s.mu.Unlock()

	user.ID = newID()
	s.Users[user.ID] = user
	return user
}
func newID() string{
	bytes := make([]byte, 8)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}