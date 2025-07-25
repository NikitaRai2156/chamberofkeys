package store

import "time"

// StringKV: operations to handle string values in the store
type StringKV interface {
	Get(key string) (string, error)
	Insert(key, val string, ttl time.Duration) error
	Update(key, val string) error
}

// ListKV: operations to handle list values in the store
type ListKV interface {
	PushFront(key, val string, ttl time.Duration) error
	PushBack(key, val string, ttl time.Duration) error
	PopFront(key string) (string, error)
	PopBack(key string) (string, error)
	GetList(key string) ([]string, error)
}

// Locks: basic concurrency control methods to control access to the store
type Locks interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

// KVStore: aggregation of all functionalities for the store
type KVStore interface {
	StringKV
	ListKV
	Locks
	Remove(key string) error
	Items() map[string]*Value
}
