package main

import (
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
)

//go:generate fyne bundle -o bundled.go ./resources/discord.png
//go:generate fyne bundle -o bundled.go -append ./resources/more.png
//go:generate fyne bundle -o bundled.go -append ./resources/mop.png

func main() {
	info_plist_modifier.SetMultipleInstancesProhibition(false)
	defer info_plist_modifier.SetMultipleInstancesProhibition(true)

	mainApp := app.New()
	window := mainApp.NewWindow("Multi Roblox Macos")
	window.Resize(fyne.NewSize(80, 80))
	window.SetFixedSize(true)

	discordButton := widget.NewButtonWithIcon("Click to join server <3", resourceDiscordPng, func() {
		discord_redirect.RedirectToServer(discord_link_parser.DiscordLink())
	})

	activateButton := widget.NewButtonWithIcon("New roblox instance", resourceMorePng, func() {
		bypass_sync.Bypass()
		<-open_app.Open("/Applications/Roblox.app")
		bypass_sync.Bypass()
	})

	closeInstancesButton := widget.NewButtonWithIcon("Close all instances", resourceMopPng, func() {
		close_all_app_instances.Close("RobloxPlayer")
	})

	window.SetContent(container.NewVBox(
		discordButton,
		activateButton,
		closeInstancesButton,
	))

	window.ShowAndRun()
}
