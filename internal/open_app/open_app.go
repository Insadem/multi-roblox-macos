package open_app

import (
	"os/exec"
)

func Open(path string) <-chan struct{} {
	var cmd string = "open"
	var args []string = []string{"-n", path}

	// open -n /Applications/Roblox.app/Contents/MacOS/RobloxPlayer

	waitChan := make(chan struct{})
	go func() {
		defer close(waitChan)
		exec.Command(cmd, args...).Run()
	}()

	return waitChan
}
