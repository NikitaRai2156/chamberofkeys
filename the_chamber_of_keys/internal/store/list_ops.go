package store

import (
	"errors"
	"time"
)

// PushFront(): to push a value to the front of a list
func (s *Store) PushFront(key, val string, ttl time.Duration) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[key]
	isExpired := s.isItemExpired(key)

	// create a new list if it does not already exist
	if !exists || isExpired {
		s.items[key] = &Value{
			Type:   ListType,
			List:   []string{val},
			Expiry: time.Now().Add(ttl),
		}
		return nil
	}

	if item.Type != ListType {
		return errors.New("operation not allowed")
	}

	// add new value to the front of the list
	item.List = append([]string{val}, item.List...)
	return nil

}

// PushBack(): to push a value to the back of a list
func (s *Store) PushBack(key, val string, ttl time.Duration) error {

	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[key]
	isExpired := s.isItemExpired(key)

	// create a new list if it does not already exist
	if !exists || isExpired {
		s.items[key] = &Value{
			Type:   ListType,
			List:   []string{val},
			Expiry: time.Now().Add(ttl),
		}
		return nil
	}

	if item.Type != ListType {
		return errors.New("operation not allowed")
	}

	// add new value to the back of the list
	item.List = append(item.List, val)
	return nil
}

// PopFront(): to pop a value from the front of a list
func (s *Store) PopFront(key string) (string, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[key]

	isExpired := s.isItemExpired(key)
	if isExpired {
		return "", errors.New("item does not exist")
	}

	if !exists || item.Type != ListType || isExpired {
		return "", errors.New("list does not exist")
	}

	if len(item.List) == 0 {
		return "", errors.New("list is empty")
	}

	// pop the first value from the list
	val := item.List[0]
	item.List = item.List[1:]

	return val, nil
}

// PopBack(): to pop a value from the back of a list
func (s *Store) PopBack(key string) (string, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[key]

	isExpired := s.isItemExpired(key)
	if isExpired {
		return "", errors.New("item does not exist")
	}

	if !exists || item.Type != ListType || isExpired {
		return "", errors.New("list does not exist")
	}

	if len(item.List) == 0 {
		return "", errors.New("list is empty")
	}

	// pop the last element from the list
	index := len(item.List) - 1
	val := item.List[index]
	item.List = item.List[:index]

	return val, nil
}

// GetList(): to get a list associated with a key
func (s *Store) GetList(key string) ([]string, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.items[key]

	isExpired := s.isItemExpired(key)
	if isExpired {
		return nil, errors.New("item does not exist")
	}

	if !exists || item.Type != ListType || isExpired {
		return nil, errors.New("list does not exist")
	}

	return item.List, nil

}
