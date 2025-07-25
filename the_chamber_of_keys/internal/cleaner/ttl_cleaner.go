package ttl

import (
	"the_chamber_of_keys/internal/store"
	"time"
)

// TTLStoreCleaner: to periodically remove expired entires from a store
type TTLStoreCleaner struct {
	Store    store.KVStore
	Interval time.Duration
	Quit     chan bool
}

// NewTTLStoreCleaner(): to create a new cleaner for the given store
func NewTTLStoreCleaner(s store.KVStore, interval time.Duration) *TTLStoreCleaner {
	return &TTLStoreCleaner{
		Store:    s,
		Interval: interval,
		Quit:     make(chan bool),
	}
}

// Start(): to launch the background cleaner to check and remove expired keys
func (c *TTLStoreCleaner) Start() {

	// run cleanup in a separate goroutine
	go func() {
		ticker := time.NewTicker(c.Interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				// time to clean expired keys
				c.cleanExpiredKeys()
			case <-c.Quit:
				// received stop signal
				return
			}
		}
	}()
}

// Stop(): to gracefully shut down the cleaner
func (c *TTLStoreCleaner) Stop() {
	close(c.Quit)
}

// cleanExpiredKeys(): to find and remove expired keys
func (c *TTLStoreCleaner) cleanExpiredKeys() {

	currentTime := time.Now()
	allItems := c.Store.Items()

	keys := make([]string, 0)

	// identify all expired keys by acquiring a read lock
	c.Store.RLock()
	for k, v := range allItems {
		if currentTime.After(v.Expiry) {
			keys = append(keys, k)
		}
	}
	c.Store.RUnlock()

	// delete all identified expired keys
	for _, k := range keys {
		c.Store.Remove(k)
	}

}
