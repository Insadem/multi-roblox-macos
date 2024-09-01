package macosapp

import (
	"github.com/ncruces/zenity"
)

func Notification(text string) {
	zenity.Notify(text, zenity.Title("multi roblox"), zenity.WarningIcon)
}
