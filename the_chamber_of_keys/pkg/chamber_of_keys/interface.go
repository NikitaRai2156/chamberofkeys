package chamberofkeys

import "time"

// ChamberOfKeys: core operations supported by the Chamber
// includes string and list-based interactions
type ChamberOfKeys interface {
	GetString(key string) (string, error)
	InsertString(key, val string, ttl time.Duration) error
	UpdateString(key, val string) error
	PushFront(key, val string, ttl time.Duration) error
	PushBack(key, val string, ttl time.Duration) error
	PopFront(key string) (string, error)
	PopBack(key string) (string, error)
	Remove(key string) error
}

// ChamberKeeper: lifecycle control methods for the ttl-based cleaner
type ChamberKeeper interface {
	StartCleaner()
	StopCleaner()
}

// ChamberManager: unified interface with key-value methods, cleaner control and system lifecycle methods
type ChamberManager interface {
	ChamberOfKeys
	ChamberKeeper
	Start()
	Stop()
}

// snapshotManager: to define internal snapshot behaviour
type snapshotManager interface {
	Restore() error
	Save() error
	StartAutoSave()
	StopAutoSave()
	SaveOnInterrupt()
}
