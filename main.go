package main

import (
	"io"
	"os/exec"

	"github.com/Insadem/multi-roblox-macos/internal/deeplink"
	"github.com/Insadem/multi-roblox-macos/internal/infoplist"
	"github.com/Insadem/multi-roblox-macos/internal/macosapp"
	"github.com/Insadem/multi-roblox-macos/internal/robloxapp"
	"github.com/Insadem/multi-roblox-macos/internal/syncbreaker"
	"github.com/Insadem/multi-roblox-macos/internal/urlhandler"
	"github.com/ncruces/zenity"
)

// Roadmap:
// 1. get rid of zenity, instead use macdriver
// 2. use embed.FS for lib folder embeding, instead of keeping it externally
// 3. rewrite deeplink package without using cgo (use purego)
func main() {
	uh, err := urlhandler.New()
	if err != nil {
		zenity.Warning("can't initialize urlhandler: " + err.Error())
		return
	}

	sb, err := syncbreaker.New()
	if err != nil {
		zenity.Warning("can't initialize syncbreaker: " + err.Error())
		return
	}

	uh.Set("com.insadem.multiroblox", "roblox-player")
	s := uh.Check(urlhandler.ROBLOX_BUNDLE_IDENTIFIER, "roblox-player")
	if s {
		zenity.Warning("can't set new urlhandler")
		return
	}

	cc := make(chan io.Closer)
	func(cc chan<- io.Closer) {
		go func() {
			for url := range deeplink.Handler() {
				copy, err := robloxapp.NewCopy()
				if err != nil {
					zenity.Notify("can't make new app copy: " + err.Error())
					continue
				}

				go func() { // not the best solution, but it's alright =). cost of 1 goroutine is ~8kb ram.
					cc <- copy
				}()

				sb.Break()
				err = exec.Command("open", "-a", copy.Path(), url).Run()
				if err != nil {
					zenity.Notify("can't open roblox app copy: " + err.Error())
					continue
				}

				err = infoplist.SetMultipleInstancesProhibition(copy.Path()+"/Contents/Info.plist", false)
				if err != nil {
					zenity.Warning("can't set multi instance prohibition: " + err.Error())
					continue
				}
				sb.Break()
			}
		}()
	}(cc)

	macosapp.Run(func() { // handle app termination here
		uh.Set(urlhandler.ROBLOX_BUNDLE_IDENTIFIER, "roblox-player")

		close(cc)
		for c := range cc {
			c.Close()
		}
	})
}
