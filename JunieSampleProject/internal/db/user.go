package db

import (
	"errors"
	"sync"
)

// User represents a user in the system
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserStore is an in-memory store for users
// In a real application, this would be replaced with a database
type UserStore struct {
	users map[string]User
	mutex sync.RWMutex
	nextID int
}

// NewUserStore creates a new user store with some initial data
func NewUserStore() *UserStore {
	store := &UserStore{
		users:  make(map[string]User),
		nextID: 3, // Start with ID 3 since we have 2 initial users
	}

	// Add some initial users
	store.users["1"] = User{ID: "1", Name: "John Doe", Email: "john@example.com"}
	store.users["2"] = User{ID: "2", Name: "Jane Smith", Email: "jane@example.com"}

	return store
}

// GetAll returns all users
func (s *UserStore) GetAll() []User {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	users := make([]User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

// GetByID returns a user by ID
func (s *UserStore) GetByID(id string) (User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}
	return user, nil
}

// Create adds a new user
func (s *UserStore) Create(name, email string) User {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	id := s.getNextID()
	user := User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	s.users[id] = user
	return user
}

// Update modifies an existing user
func (s *UserStore) Update(id, name, email string) (User, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("user not found")
	}

	// Update fields if provided
	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}

	s.users[id] = user
	return user, nil
}

// Delete removes a user
func (s *UserStore) Delete(id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, exists := s.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(s.users, id)
	return nil
}

// getNextID returns the next available ID as a string
func (s *UserStore) getNextID() string {
	id := s.nextID
	s.nextID++
	return string(id + '0') // Convert to string
}