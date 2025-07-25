package persistence

import (
	"the_chamber_of_keys/internal/store"
	"time"
)

// SerializedValue: a flat format of store.Value for persistent storage
type SerializedValue struct {
	Key    string
	Type   store.DataType
	String string
	List   []string
	Expiry int64 // Unix timestamp
}

// Serialize(): to convert in-memory format store.Value into database-friendly SerializedValue
func Serialize(store store.KVStore) ([]SerializedValue, error) {
	store.RLock()
	defer store.RUnlock()

	var serialized []SerializedValue
	for key, val := range store.Items() {
		serialized = append(serialized, SerializedValue{
			Key:    key,
			Type:   val.Type,
			String: val.String,
			List:   val.List,
			Expiry: val.Expiry.Unix(),
		})
	}
	return serialized, nil
}

// Deserialize(): to reconstruct serialized values in the database into im-memory usable store.Value
func Deserialize(data []SerializedValue) map[string]*store.Value {
	result := make(map[string]*store.Value)
	for _, record := range data {
		result[record.Key] = &store.Value{
			Type:   record.Type,
			String: record.String,
			List:   record.List,
			Expiry: time.Unix(record.Expiry, 0),
		}
	}
	return result
}
