package fspath

import "testing"

func TestTMPDir(t *testing.T) {
	pth, err := TMPDir.Get()
	if err != nil {
		t.Error(err)
	}

	if pth == "" {
		t.Error("expected non empty path")
	}
}
