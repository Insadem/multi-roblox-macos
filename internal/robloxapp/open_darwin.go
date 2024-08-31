//go:build darwin
// +build darwin

package robloxapp

import "os/exec"

func Open() (func(), error) {
	cmd := exec.Command("/Applications/Roblox.app/Contents/MacOS/RobloxPlayer")
	err := cmd.Start()
	if err != nil {
		return nil, err
	}

	return func() {
		cmd.Process.Kill()
		cmd.Process.Release()
	}, nil
}
