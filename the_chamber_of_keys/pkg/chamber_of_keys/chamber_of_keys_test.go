package chamberofkeys

import (
	"testing"
	"the_chamber_of_keys/internal/store"
	"time"
)

func TestChamber_InsertString(t *testing.T) {

	testChamber := Chamber{
		kvstore: &store.MockStore{},
	}

	type args struct {
		key string
		val string
		ttl time.Duration
	}

	tests := []struct {
		name    string
		chamber ChamberOfKeys
		args    args
		wantErr bool
	}{
		{
			name:    "Case 1: Successful Insert",
			chamber: &testChamber,
			args:    args{key: "10", val: "abc", ttl: 5 * time.Minute},
			wantErr: false,
		},
		{
			name:    "Case 2: Inserting key that already exists",
			chamber: &testChamber,
			args:    args{key: "1", val: "abc", ttl: 5 * time.Minute},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.chamber.InsertString(tt.args.key, tt.args.val, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("Chamber.InsertString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
