package shortenpath

import (
	"testing"
)

type errInvalidPath struct {
	path string
}

func (err errInvalidPath) Error() string {
	return "invalid path: " + err.path
}

type errInvalidVal struct {
	val string
}

func (err errInvalidVal) Error() string {
	return "invalid val: " + err.val
}

func TestShortenPath(t *testing.T) {
	pth, v := Shorten("usr/bin/cake")
	if pth != "usr/bin" {
		t.Error(errInvalidPath{path: pth})
		return
	}
	if v != "cake" {
		t.Error(errInvalidVal{val: v})
		return
	}

	pth, v = Shorten(pth)
	if pth != "usr" {
		t.Error(errInvalidPath{path: pth})
		return
	}
	if v != "bin" {
		t.Error(errInvalidVal{val: v})
		return
	}

	pth, v = Shorten(pth)
	if pth != "" {
		t.Error(errInvalidPath{path: pth})
		return
	}
	if v != "usr" {
		t.Error(errInvalidVal{val: v})
		return
	}

	pth, v = Shorten(pth)
	if pth != "" || v != "" {
		t.Error("expected to be clean")
	}
}
