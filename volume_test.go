package netapp

import "testing"

func TestGetVolumesV2(t *testing.T) {
	got,err := GetVolumesV2("ep9")
	if err != nil {
		t.Error(err)
		return
	}
	if len(got.Records) == 0 {
		t.Error("nothing found")
	}
}