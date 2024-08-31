package fspath

import "testing"

func TestFindLibDir(t *testing.T) {
	dir := findLibDir("usr/multi-roblox-macos/lib/keep/search/for/the/truth")
	if dir != "usr/multi-roblox-macos/lib" {
		t.Error("didn't found lib dir, instead found: " + dir)
	}

	dir = findLibDir("usr/MacOS/multi-roblox-macos")
	if dir != "usr/MacOS/lib" {
		t.Error("expected to find lib via MacOS folder")
	}

	dir = findLibDir("there/is/no/truth")
	if dir != "" {
		t.Error("expected to be empty")
	}
}

func TestLibDir(t *testing.T) {
	_, err := LibDir.Get()
	if err != nil {
		t.Error(err)
	}
}
