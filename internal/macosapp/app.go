package macosapp

import (
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/macos/webkit"
	"github.com/progrium/darwinkit/objc"
)

func Run(clean func()) {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		url := foundation.URL_URLWithString("https://discord.gg/AwNAa7utbY")
		req := foundation.NewURLRequestWithURL(url)
		frame := foundation.Rect{Size: foundation.Size{Width: 480, Height: 320}}

		config := webkit.NewWebViewConfiguration()
		wv := webkit.NewWebViewWithFrameConfiguration(frame, config)
		wv.LoadRequest(req)

		w := appkit.NewWindowWithContentRectStyleMaskBackingDefer(frame,
			appkit.ClosableWindowMask|appkit.TitledWindowMask,
			appkit.BackingStoreBuffered, false)
		objc.Retain(&w)
		w.SetTitle("multi-roblox-macos")
		w.SetContentView(wv)
		w.MakeKeyAndOrderFront(w)
		w.Center()
		w.SetStyleMask(w.StyleMask() | appkit.MiniaturizableWindowMask)

		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			clean()
			return true
		})
	})
}
