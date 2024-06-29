package backup_roblox_app

import "testing"

func TestBackup(t *testing.T) {
	result := <-NewBackup()
	if result.Err != nil {
		t.Error(result.Err)
	}
}

func TestClear(t *testing.T) {
	err := ClearBackup()
	if err != nil {
		t.Error(err)
	}
}
