package store

import (
	"errors"
	"sync"
	"time"
)

// DataType: to define possible types of data in the store
type DataType int

const (
	StringType DataType = iota
	ListType
)

// Value: to store actual data with its type, value, and expiry
type Value struct {
	Type   DataType
	String string
	List   []string
	Expiry time.Time
}

// Store: the main in-memory store
type Store struct {
	mu    sync.RWMutex
	items map[string]*Value
}

// NewStore(): to create a new instance of Store
func NewStore() *Store {
	return &Store{
		mu:    sync.RWMutex{},
		items: make(map[string]*Value),
	}
}

// Remove(): to remove a key from the store if it exists
func (s *Store) Remove(key string) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.items[key]

	if !exists {
		return errors.New("invalid key")
	}

	delete(s.items, key)
	return nil
}

// Lock(): to acquire a write lock on the store
func (s *Store) Lock() {
	s.mu.Lock()
}

// Unlock(): to release the write lock on the store
func (s *Store) Unlock() {
	s.mu.Unlock()
}

// RLock(): to acquire a read lock on the store
func (s *Store) RLock() {
	s.mu.RLock()
}

// RUnlock(): to release the read lock on the store
func (s *Store) RUnlock() {
	s.mu.RUnlock()
}

// Items(): to return all the items in the store
func (s *Store) Items() map[string]*Value {
	return s.items
}

// isItemExpired: to check if a given key has expired
// returns true if the key does not exist or has expired
func (s *Store) isItemExpired(key string) bool {

	item, exists := s.items[key]

	if !exists {
		return true
	}

	if time.Now().After(item.Expiry) {
		return true
	}

	return false
}
