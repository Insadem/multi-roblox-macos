package bypass_sync

import "testing"

func TestBypass(t *testing.T) {
	err := Bypass()
	if err != nil {
		t.Error(err)
	}
}
