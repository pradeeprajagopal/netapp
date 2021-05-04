package netapp

import (
	"reflect"
	"testing"
)

func TestGetAllNfsInfo(t *testing.T) {
	tests := []struct {
		name    string
		want    NFS
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAllNfsInfo()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAllNfsInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllNfsInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
