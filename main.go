package main

import (
	"insadem/multi_roblox_macos/internal/bypass_sync"
	"insadem/multi_roblox_macos/internal/discord_link_parser"
	"insadem/multi_roblox_macos/internal/discord_redirect"
	"insadem/multi_roblox_macos/internal/open_app"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	mainApp := app.New()
	window := mainApp.NewWindow("Multi Roblox Macos")
	window.Resize(fyne.NewSize(80, 80))
	window.SetFixedSize(true)

	discordButton := widget.NewButtonWithIcon("Click to join server <3.", resourceDiscordPng, func() {
		discord_redirect.RedirectToServer(discord_link_parser.DiscordLink())
	})

	var activateButton *widget.Button
	activateButton = widget.NewButtonWithIcon("Add roblox instance", resourceMorePng, func() { // Set to active / not active
		bypass_sync.BypassSync() // Isn't necessary, OpenApp bellow does all job, but still let's keep it as safeguard.
		open_app.OpenApp("/Applications/Roblox.app")
	})

	window.SetContent(container.NewVBox(
		discordButton,
		activateButton,
	))

	window.ShowAndRun()
}
