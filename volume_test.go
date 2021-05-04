package netapp

import (
	"reflect"
	"testing"
)

func TestGetVolumesV2(t *testing.T) {
	tests := []struct {
		name    string
		want    VolumeV2
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVolumesV2()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVolumesV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVolumesV2() got = %v, want %v", got, tt.want)
			}
		})
	}
}
