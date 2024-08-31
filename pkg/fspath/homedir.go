package fspath

import "os"

var HomeDir = New(func() (string, error) {
	return os.UserHomeDir()
})
