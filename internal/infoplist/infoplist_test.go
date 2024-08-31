package infoplist

import "testing"

func TestSetMultipleInstancesProhibition(t *testing.T) {
	err := SetMultipleInstancesProhibition("/Applications/Roblox.app/Contents/Info.plist", false)
	if err != nil {
		t.Error(err)
	}
}
