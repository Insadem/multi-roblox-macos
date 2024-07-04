package backup_roblox_app

import (
	"errors"
	"os"
	"os/exec"
)

func getBackupDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/roblox_backup", nil
}

func getDestinationDir() (string, error) {
	backupDir, err := getBackupDir()
	if err != nil {
		return "", err
	}

	err = os.Mkdir(backupDir, 0777)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return "", err
	}

	path, err := os.MkdirTemp(backupDir, "")
	return path, err
}

type BackupResult struct {
	Path string
	Err  error
}

func NewBackup() <-chan BackupResult {
	resultChan := make(chan BackupResult)

	go func() {
		defer close(resultChan)

		srcPath := "/Applications/Roblox.app"
		destDir, err := getDestinationDir()
		if err != nil {
			resultChan <- BackupResult{Err: err}
			return
		}

		cpCmd := exec.Command("cp", "-a", srcPath, destDir)
		err = cpCmd.Run()
		resultChan <- BackupResult{Path: destDir + "/Roblox.app", Err: err}
	}()

	return resultChan
}

func ClearBackup() error {
	backupDir, err := getBackupDir()
	if err != nil {
		return err
	}

	// TODO: sometimes might error, so retry mechanism at least for 1-2 times is quite good solution?
	return os.RemoveAll(backupDir)
}
