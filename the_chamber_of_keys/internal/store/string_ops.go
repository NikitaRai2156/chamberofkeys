package store

import (
	"errors"
	"time"
)

// Get(): to get the string value associated with the key
func (s *Store) Get(key string) (string, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[key]

	isExpired := s.isItemExpired(key)

	if !exists || item.Type != StringType || isExpired {
		return "", errors.New("string does not exist")
	}

	return item.String, nil
}

// Insert(): to insert a new string value with a given key
func (s *Store) Insert(key, val string, ttl time.Duration) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.items[key]

	if exists && !s.isItemExpired(key) {
		return errors.New("string already exists")
	}

	s.items[key] = &Value{
		Type:   StringType,
		String: val,
		Expiry: time.Now().Add(ttl),
	}

	return nil
}

// Update(): to update the string value of an existing key
func (s *Store) Update(key, val string) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[key]

	isExpired := s.isItemExpired(key)

	if !exists || item.Type != StringType || isExpired {
		return errors.New("invalid key")
	}

	item.String = val
	return nil
}
