package store

import (
	"sync"
	"testing"
	"time"
)

var test_values_list = []*Value{
	{Type: ListType, List: []string{"f", "g"}, Expiry: setExpiryTime(5 * time.Minute)},
	{Type: ListType, List: []string{"h", "i"}, Expiry: setExpiryTime(5 * time.Minute)},
	{Type: ListType, List: []string{"j", "k"}, Expiry: setExpiryTime(5 * time.Minute)},
}

var test_store_list = Store{
	mu: sync.RWMutex{},
	items: map[string]*Value{
		"1": test_values_list[0],
		"2": test_values_list[1],
		"3": test_values_list[2],
	},
}

func TestStore_PushFront(t *testing.T) {

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
			name:    "Case 1: Succesful push to front when list exists",
			store:   &test_store_list,
			args:    args{key: "1", val: "z", ttl: 5 * time.Minute},
			wantErr: false,
		},
		{
			name:    "Case 2: Succesful push to front when list does not exist",
			store:   &test_store_list,
			args:    args{key: "4", val: "z", ttl: 5 * time.Minute},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.store.PushFront(tt.args.key, tt.args.val, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("Store.PushFront() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
