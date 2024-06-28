package open_app

import (
	"os/exec"
)

func OpenApp(path string) error {
	var cmd string = "open"
	var args []string = []string{"-n", path}

	// open -n /Applications/Roblox.app/Contents/MacOS/RobloxPlayer

	cmdInstance := exec.Command(cmd, args...)
	return cmdInstance.Start()
}
