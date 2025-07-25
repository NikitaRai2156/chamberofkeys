package snapshot

type Manager interface {
	Restore() error
	Save() error
	StartAutoSave()
	StopAutoSave()
	SaveOnInterrupt()
}
