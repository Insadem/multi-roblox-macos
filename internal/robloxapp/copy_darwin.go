package robloxapp

import "os/exec"

func NewCopy() (Copy, error) {
	p := "/Applications/Roblox.app"
	d, err := copyDestDir()
	if err != nil {
		return Copy{}, err
	}

	err = exec.Command("cp", "-a", p, d).Run()
	if err != nil {
		return Copy{}, err
	}

	return Copy{dir: d}, nil
}
