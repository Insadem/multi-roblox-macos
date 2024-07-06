package new_instance_widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Insadem/multi-roblox-macos/internal/backup_roblox_app"
	"github.com/Insadem/multi-roblox-macos/internal/bypass_sync"
	"github.com/Insadem/multi-roblox-macos/internal/info_plist_modifier"
	"github.com/Insadem/multi-roblox-macos/internal/open_app"
	"github.com/ncruces/zenity"
)

// it does also bypass Mac signature check.
func initiateMIPBypass(path string) error {
	err := <-open_app.Open(path)
	if err != nil {
		return err
	}
	err = info_plist_modifier.SetMultipleInstancesProhibition(path+"/Contents/Info.plist", false)
	if err != nil {
		return err
	}

	return nil
}

// Getter
func NewInstanceButton(icon fyne.Resource) *widget.Button {
	var newInstanceButton *widget.Button
	newInstanceButton = widget.NewButtonWithIcon("New roblox instance", icon, func() {
		if newInstanceButton.Disabled() {
			return
		}
		newInstanceButton.Disable()

		go func() {
			defer newInstanceButton.Enable()

			progressDlg, err := zenity.Progress(
				zenity.Title("Creating new roblox instance."),
				zenity.Pulsate())
			if err != nil {
				return
			}
			defer progressDlg.Close()

			backupResult := <-backup_roblox_app.NewBackup()
			if backupResult.Err != nil {
				zenity.Error("Couldn't create roblox app backup. Does roblox app exist?")
				return
			}

			err = initiateMIPBypass(backupResult.Path)
			if err != nil {
				zenity.Error("Couldn't iniate MIP bypass. Try click again.")
				return
			}

			bypass_sync.Bypass()
			<-open_app.Open(backupResult.Path)
			bypass_sync.Bypass()
		}() // new goroutine to prevent main UI thread block
	})
	return newInstanceButton
}
