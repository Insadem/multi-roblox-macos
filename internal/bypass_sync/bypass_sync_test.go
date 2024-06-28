package bypass_sync

import "testing"

func TestBypassSync(t *testing.T) {
	err := BypassSync()
	if err != nil {
		t.Error(err)
	}
}
