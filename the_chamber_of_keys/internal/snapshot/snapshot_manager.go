package snapshot

import (
	"os"
	"os/signal"
	"time"

	"the_chamber_of_keys/internal/persistence"
	"the_chamber_of_keys/internal/store"
)

// SnapshotManager: to handle the periodic saving and restoration of the in-memory store
type SnapshotManager struct {
	store      store.KVStore
	persist    persistence.PStore
	interval   time.Duration
	stopTicker chan struct{}
}

// New(): to create a new SnapshotManager with the given store, persistence and interval
func New(store store.KVStore, persist persistence.PStore, interval time.Duration) *SnapshotManager {
	return &SnapshotManager{
		store:      store,
		persist:    persist,
		interval:   interval,
		stopTicker: make(chan struct{}),
	}
}

// Restore(): load the last saved snapshot and replace current in-memory store
func (s *SnapshotManager) Restore() error {
	data, err := s.persist.Load()
	if err != nil {
		return err
	}
	items := persistence.Deserialize(data)

	s.store.Lock()
	defer s.store.Unlock()
	for k := range s.store.Items() {
		delete(s.store.Items(), k)
	}
	for k, v := range items {
		s.store.Items()[k] = v
	}
	return nil
}

// Save(): to create and save a snapshot of the current in-memory store
func (s *SnapshotManager) Save() error {
	data, err := persistence.Serialize(s.store)
	if err != nil {
		return err
	}
	return s.persist.Save(data)
}

// StartAutoSave(): to start saving the in-memory store at regular intervals to the persistent db
func (s *SnapshotManager) StartAutoSave() {
	ticker := time.NewTicker(s.interval)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.Save()
			case <-s.stopTicker:
				ticker.Stop()
				return
			}
		}
	}()
}

// StopAutoSave(): to signal the background auto-save routine to stop
func (s *SnapshotManager) StopAutoSave() {
	close(s.stopTicker)
}

// SaveOnInterrupt(): to save the in-memory store on an interrupt signal
func (s *SnapshotManager) SaveOnInterrupt() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)
	go func() {
		<-sigchan
		s.Save()
		os.Exit(0)
	}()
}
