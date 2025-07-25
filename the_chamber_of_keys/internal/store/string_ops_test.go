package store

import (
	"sync"
	"testing"
	"time"
)

func setExpiryTime(t time.Duration) time.Time {
	return time.Now().Add(t)
}

var test_values_string = []*Value{
	{Type: StringType, String: "a", Expiry: setExpiryTime(5 * time.Minute)},
	{Type: StringType, String: "b", Expiry: setExpiryTime(5 * time.Minute)},
	{Type: StringType, String: "c", Expiry: setExpiryTime(5 * time.Minute)},
	{Type: StringType, String: "d", Expiry: setExpiryTime(5 * time.Minute)},
}

var test_store_string = Store{
	mu: sync.RWMutex{},
	items: map[string]*Value{
		"1": test_values_string[0],
		"2": test_values_string[1],
		"3": test_values_string[2],
		"4": test_values_string[3],
	},
}

func TestStore_Insert(t *testing.T) {

	type args struct {
		key string
		val string
		ttl time.Duration
	}

	tests := []struct {
		name    string
		store   *Store
		args    args
		wantErr bool
	}{
		{
			name:    "Case 1: Successful insert",
			store:   &test_store_string,
			args:    args{key: "7", val: "d", ttl: 5 * time.Minute},
			wantErr: false,
		},
		{
			name:    "Case 2: Inserting key that already exists",
			store:   &test_store_string,
			args:    args{key: "1", val: "d", ttl: 5 * time.Minute},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.store.Insert(tt.args.key, tt.args.val, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("Store.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
