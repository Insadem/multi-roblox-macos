package main

import (
	"insadem/multi_roblox_macos/internal/backup_roblox_app"
	"insadem/multi_roblox_macos/internal/close_all_app_instances"
	"insadem/multi_roblox_macos/internal/discord_link_parser"
	"insadem/multi_roblox_macos/internal/discord_redirect"
	new_instance_widget "insadem/multi_roblox_macos/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//go:generate fyne bundle -o bundled.go ./resources/discord.png
//go:generate fyne bundle -o bundled.go -append ./resources/more.png
//go:generate fyne bundle -o bundled.go -append ./resources/mop.png

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

	closeInstancesButton := widget.NewButtonWithIcon("Close all instances", resourceMopPng, func() {
		close_all_app_instances.Close("RobloxPlayer")
	})

	window.SetContent(container.NewVBox(
		discordButton,
		new_instance_widget.NewInstanceButton(resourceMorePng),
		closeInstancesButton,
	))

	window.ShowAndRun()
}
