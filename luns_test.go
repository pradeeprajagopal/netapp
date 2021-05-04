package netapp

import (
	"reflect"
	"testing"
)

func TestGetLuns(t *testing.T) {
	tests := []struct {
		name    string
		want    Luns
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetLuns()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLuns() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLuns() got = %v, want %v", got, tt.want)
			}
		})
	}
}
