package chamberofkeys

import (
	"fmt"
	"os"
	"path/filepath"
	ttl "the_chamber_of_keys/internal/cleaner"
	"the_chamber_of_keys/internal/persistence"
	"the_chamber_of_keys/internal/snapshot"
	"the_chamber_of_keys/internal/store"
	"time"
)

// Chamber: core struct - bundles the key-value store, ttl-based cleaner, and snapshot manager
type Chamber struct {
	kvstore store.KVStore
	cleaner ttl.Cleaner
	manager snapshotManager
}

// getProjectDataPath(): set up the file path for the snapshot database
func getProjectDataPath() string {
	dir := "data"
	_ = os.MkdirAll(dir, 0755)
	return filepath.Join(dir, "snapshot.db")
}

// NewChamber(): sets up and returns a fully initialised Chamber instance
func NewChamber() (*Chamber, error) {

	newstore := store.NewStore()

	// new cleaner to clean up expired keys from the store every 1 hour
	newcleaner := ttl.NewTTLStoreCleaner(newstore, time.Hour)

	dbpath := getProjectDataPath()
	persist, err := persistence.NewSQLiteStore(dbpath)
	if err != nil {
		fmt.Printf("Error while initialising persistence in chamber: %v", err.Error())
		return nil, err
	}

	// snapshot manager to take a snapshot of the system every 5 minutes
	snapman := snapshot.New(newstore, persist, 5*time.Minute)

	return &Chamber{
		kvstore: newstore,
		cleaner: newcleaner,
		manager: snapman,
	}, nil
}

// Start(): launches the Chamber services
// including ttl-based cleaner, restoration from disk, and periodic auto-saving of a snapshot of the store
func (c *Chamber) Start() {
	c.StartCleaner()
	c.manager.Restore()
	c.manager.StartAutoSave()
	c.manager.SaveOnInterrupt()
}

// Stop(): to safely shut down all Chamber services
func (c *Chamber) Stop() {
	c.StopCleaner()
	c.manager.StopAutoSave()
	c.manager.Save()
}

// GetString(): to retrieve a string value by key
func (c *Chamber) GetString(key string) (string, error) {
	return c.kvstore.Get(key)
}

// InsertString(): to add a string value in the store
func (c *Chamber) InsertString(key, val string, ttl time.Duration) error {
	return c.kvstore.Insert(key, val, ttl)
}

// UpdateString(): to modify the string value for an existing key
func (c *Chamber) UpdateString(key, val string) error {
	return c.kvstore.Update(key, val)
}

// PushFront(): to add an element to the front of a list value
func (c *Chamber) PushFront(key, val string, ttl time.Duration) error {
	return c.kvstore.PushFront(key, val, ttl)
}

// PushBack(): to add an element to the back of a list value
func (c *Chamber) PushBack(key, val string, ttl time.Duration) error {
	return c.kvstore.PushBack(key, val, ttl)
}

// PopFront(): to remove and return the front element of a list
func (c *Chamber) PopFront(key string) (string, error) {
	return c.kvstore.PopFront(key)
}

// PopBack(): to remove and return the last element of a list
func (c *Chamber) PopBack(key string) (string, error) {
	return c.kvstore.PopBack(key)
}

// Remove(): to delete a key and its associated value from the store
func (c *Chamber) Remove(key string) error {
	return c.kvstore.Remove(key)
}

// StartCleaner(): to launch the cleaner
func (c *Chamber) StartCleaner() {
	c.cleaner.Start()
}

// StopCleaner(): to stop the cleaner
func (c *Chamber) StopCleaner() {
	c.cleaner.Stop()
}
