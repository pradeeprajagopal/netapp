package netapp

import (
	"reflect"
	"testing"
)

func TestGetInterfaces(t *testing.T) {
	tests := []struct {
		name    string
		want    Interfaces
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInterfaces()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInterfaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInterfaces() got = %v, want %v", got, tt.want)
			}
		})
	}
}
