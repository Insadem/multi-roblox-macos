package fspath

import (
	"errors"
	"testing"
)

func TestPath_Success(t *testing.T) {
	path := New(func() (string, error) {
		return "/some/path/", nil
	})

	pth, err := path.Get()
	if err != nil || pth != "/some/path/" {
		t.Error("expected success")
		return
	}

	path.Set("/some/new/path/")
	pth, err = path.Get()
	if err != nil || pth != "/some/new/path/" {
		t.Error("expected new path success")
		return
	}
}

func TestPath_Bad(t *testing.T) {
	path := New(func() (string, error) {
		return "", errors.New("some error")
	})

	_, err := path.Get()
	if err == nil {
		t.Error("expected bad")
		return
	}
}
