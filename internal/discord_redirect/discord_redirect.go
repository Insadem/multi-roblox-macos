package discord_redirect

import (
	"os/exec"
)

func openInDefaultBrowser(url string) error {
	var cmd string = "open"
	var args []string

	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func RedirectToServer(inviteLink string) {
	openInDefaultBrowser(inviteLink)
}
