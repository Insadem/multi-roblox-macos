package open_app

import (
	"os/exec"
)

func Open(path string) <-chan error {
	var cmd string = "open"
	var args []string = []string{"-n", path}

	// open -n /Applications/Roblox.app/Contents/MacOS/RobloxPlayer

	resultChan := make(chan error)
	go func() {
		defer close(resultChan)
		resultChan <- exec.Command(cmd, args...).Run()
	}()

	return resultChan
}
