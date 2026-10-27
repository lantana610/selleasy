package models

func (s *Store) CreateSession(userID string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	token := newID() + newID()
	s.Sessions[token] = userID
	return token
}
func (s *Store) GetOrder(id string) (Order, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	order, ok := s.Orders[id]
	return order, ok
}

func (s *Store) SessionUser(token string) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	userID, ok := s.Sessions[token]
	if !ok {
		return User{}, false
	}
	user, ok := s.Users[userID]
	return user, ok
}

func (s *Store) DeleteSession(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.Sessions, token)
}
