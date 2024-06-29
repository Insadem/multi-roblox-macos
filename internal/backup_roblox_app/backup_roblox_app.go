package backup_roblox_app

import (
	"errors"
	"os"
	"os/exec"
)

func getDestinationFolder() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/roblox_backup", nil
}

type BackupResult struct {
	Path string
	Err  error
}

func NewBackup() <-chan BackupResult {
	resultChan := make(chan BackupResult)

	go func() {
		defer close(resultChan)

		srcFolder := "/Applications/Roblox.app"
		destFolder, err := getDestinationFolder()
		if err != nil {
			resultChan <- BackupResult{Err: err}
			return
		}

		err = os.Mkdir(destFolder, 0777)
		if err != nil && !errors.Is(err, os.ErrExist) {
			resultChan <- BackupResult{Err: err}
			return
		}

		cpCmd := exec.Command("cp", "-rf", srcFolder, destFolder)
		err = cpCmd.Run()
		resultChan <- BackupResult{Path: destFolder + "/Roblox.app", Err: err}
	}()

	return resultChan
}

func ClearBackup() error {
	destFolder, err := getDestinationFolder()
	if err != nil {
		return err
	}

	return os.RemoveAll(destFolder)
}
