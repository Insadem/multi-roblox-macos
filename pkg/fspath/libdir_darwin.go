package fspath

import (
	"os"

	"github.com/Insadem/multi-roblox-macos/pkg/fspath/shortenpath"
)

// findLibDir returns path to lib dir or empty string.
func findLibDir(path string) string {
	for {
		pth, v := shortenpath.Shorten(path)
		if pth == "" {
			return ""
		}
		if v == "multi-roblox-macos" {
			_, v = shortenpath.Shorten(pth)
			if v == "MacOS" {
				return pth + "/lib" // load in prod
			}

			return pth + "/multi-roblox-macos/lib"
		}

		path = pth
	}
}

func libDirPath(path string) (string, error) {
	ld := findLibDir(path)
	if ld == "" {
		return "", os.ErrNotExist
	}

	dir := "/" + ld
	if _, err := os.Stat(dir); err != nil && os.IsNotExist(err) {
		return "", err
	}

	return dir, nil
}

var LibDir = New(func() (string, error) {
	pth, err := os.Executable()
	if err != nil {
		return "", err
	}

	pth, err = libDirPath(pth)
	if err != nil {
		// Try fallback to current work directory, for example when testing
		pth, err = os.Getwd()
		if err != nil {
			return "", err
		}

		pth, err = libDirPath(pth)
		if err != nil {
			return "", err
		}

		return pth, nil
	}

	return pth, nil
})
