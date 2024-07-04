package main

import (
	"insadem/multi_roblox_macos/internal/backup_roblox_app"
	"insadem/multi_roblox_macos/internal/bypass_sync"
	"insadem/multi_roblox_macos/internal/close_all_app_instances"
	"insadem/multi_roblox_macos/internal/discord_link_parser"
	"insadem/multi_roblox_macos/internal/discord_redirect"
	"insadem/multi_roblox_macos/internal/info_plist_modifier"
	"insadem/multi_roblox_macos/internal/open_app"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/ncruces/zenity"
)

//go:generate fyne bundle -o bundled.go ./resources/discord.png
//go:generate fyne bundle -o bundled.go -append ./resources/more.png
//go:generate fyne bundle -o bundled.go -append ./resources/mop.png

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

func main() {
	backup_roblox_app.ClearBackup()
	defer backup_roblox_app.ClearBackup()

	close_all_app_instances.Close("RobloxPlayer")

	mainApp := app.New()
	window := mainApp.NewWindow("Multi Roblox Macos")
	window.Resize(fyne.NewSize(80, 80))
	window.SetFixedSize(true)

	discordButton := widget.NewButtonWithIcon("Click to join server <3", resourceDiscordPng, func() {
		discord_redirect.RedirectToServer(discord_link_parser.DiscordLink())
	})

	// add progress bar that prevents user from clicking another button

	var newInstanceButton *widget.Button
	newInstanceButton = widget.NewButtonWithIcon("New roblox instance", resourceMorePng, func() {
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

	closeInstancesButton := widget.NewButtonWithIcon("Close all instances", resourceMopPng, func() {
		close_all_app_instances.Close("RobloxPlayer")
	})

	window.SetContent(container.NewVBox(
		discordButton,
		newInstanceButton,
		closeInstancesButton,
	))

	window.ShowAndRun()
}
