package robloxapp

import (
	"errors"
	"io"
	"os"

	"github.com/Insadem/multi-roblox-macos/pkg/fspath"
)

type Copy struct {
	dir string
}

func (b Copy) Path() string {
	return b.dir + "/Roblox.app"
}

// Close releases all associated resources with the copy.
func (b Copy) Close() error {
	return os.RemoveAll(b.dir)
}

var _ io.Closer = (*Copy)(nil)

func copyDestDir() (string, error) {
	d, err := fspath.TMPDir.Get()
	if err != nil {
		return "", err
	}

	err = os.Mkdir(d, 0777)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return "", err
	}

	path, err := os.MkdirTemp(d, "")
	return path, err
}
