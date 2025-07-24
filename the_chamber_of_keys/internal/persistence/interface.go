package persistence

type PStore interface {
	Save(data []SerializedValue) error
	Load() ([]SerializedValue, error)
}
