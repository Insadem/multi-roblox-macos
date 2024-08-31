package fspath

import (
	"errors"
	"os"
)

var TMPDir = New(func() (string, error) {
	d := os.TempDir()
	if d == "" {
		return "", errors.New("unable to locate temp dir")
	}

	return d, nil
})
