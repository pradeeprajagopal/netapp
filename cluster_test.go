package netapp

import (
	"reflect"
	"testing"
)

func TestGetClusters(t *testing.T) {
	tests := []struct {
		name    string
		want    Clusters
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetClusters()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetClusters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetClusters() got = %v, want %v", got, tt.want)
			}
		})
	}
}
