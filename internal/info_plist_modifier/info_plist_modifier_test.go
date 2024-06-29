package info_plist_modifier

import "testing"

func TestSetMultipleInstancesProhibition(t *testing.T) {
	err := SetMultipleInstancesProhibition(false)
	if err != nil {
		t.Error(err)
	}
}
