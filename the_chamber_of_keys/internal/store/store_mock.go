package store

import (
	"errors"
	"time"
)

const mockValue = "abc"

func setMockExpiryTime(t time.Duration) time.Time {
	return time.Now().Add(t)
}

var mock_store = map[string]*Value{
	"1": {Type: StringType, String: "a", Expiry: setMockExpiryTime(5 * time.Minute)},
	"2": {Type: StringType, String: "b", Expiry: setMockExpiryTime(5 * time.Minute)},
	"3": {Type: StringType, String: "c", Expiry: setMockExpiryTime(5 * time.Minute)},
	"4": {Type: StringType, String: "d", Expiry: setMockExpiryTime(5 * time.Minute)},
	"5": {Type: ListType, List: []string{"f", "g"}, Expiry: setMockExpiryTime(5 * time.Minute)},
	"6": {Type: ListType, List: []string{"h", "i"}, Expiry: setMockExpiryTime(5 * time.Minute)},
	"7": {Type: ListType, List: []string{"j", "k"}, Expiry: setMockExpiryTime(5 * time.Minute)},
}

type MockStore struct{}

func (s *MockStore) Get(key string) (string, error) { return mockValue, nil }

func (s *MockStore) Insert(key, val string, ttl time.Duration) error {
	if _, exists := mock_store[key]; exists {
		return errors.New("key already exists")
	}
	return nil
}

func (s *MockStore) Update(key, val string) error { return nil }

func (s *MockStore) PushFront(key, val string, ttl time.Duration) error { return nil }

func (s *MockStore) PushBack(key, val string, ttl time.Duration) error { return nil }

func (s *MockStore) PopFront(key string) (string, error) { return mockValue, nil }

func (s *MockStore) PopBack(key string) (string, error) { return mockValue, nil }

func (s *MockStore) GetList(key string) ([]string, error) { return []string{mockValue}, nil }

func (s *MockStore) Lock() {}

func (s *MockStore) Unlock() {}

func (s *MockStore) RLock() {}

func (s *MockStore) RUnlock() {}

func (s *MockStore) Remove(key string) error { return nil }

func (s *MockStore) Items() map[string]*Value {
	return mock_store
}
