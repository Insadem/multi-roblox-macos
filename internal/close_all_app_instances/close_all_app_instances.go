package close_all_app_instances

import (
	"os"

	"github.com/Insadem/multi-roblox-macos/internal/ps_darwin"
)

func Close(name string) {
	processes, err := ps_darwin.Processes()
	if err != nil {
		return
	}

	for _, process := range processes {
		pName := process.Executable()
		if pName == name {
			p, err := os.FindProcess(process.Pid())
			if err == nil {
				p.Kill()
			}
		}
	}
}
