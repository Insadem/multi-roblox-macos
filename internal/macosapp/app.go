package macosapp

import (
	"os/exec"

	"github.com/Insadem/multi-roblox-macos/internal/robloxapp"
	"github.com/progrium/darwinkit/helper/action"
	"github.com/progrium/darwinkit/helper/layout"
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
)

func Run(clean func()) {
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		w := appkit.NewWindowWithContentRectStyleMaskBackingDefer(foundation.Rect{Size: foundation.Size{Width: 32, Height: 32}},
			appkit.ClosableWindowMask|appkit.TitledWindowMask|appkit.MiniaturizableWindowMask,
			appkit.BackingStoreBuffered, false)
		objc.Retain(&w)

		lable := appkit.NewLabel("start any roblox game via browser")
		dButton := appkit.NewButtonWithTitle("join discord server")
		cButton := appkit.NewButtonWithTitle("close all instances")

		action.Set(dButton, func(sender objc.Object) {
			exec.Command("open", "https://discord.gg/AwNAa7utbY").Run()
		})
		action.Set(cButton, func(sender objc.Object) {
			robloxapp.CloseAll()
		})

		stackView := appkit.StackView_StackViewWithViews([]appkit.IView{lable, dButton, cButton})
		stackView.SetOrientation(appkit.UserInterfaceLayoutOrientationVertical)
		stackView.SetDistribution(appkit.StackViewDistributionFillEqually)
		stackView.SetAlignment(appkit.LayoutAttributeCenterX)
		stackView.SetSpacing(10)

		w.ContentView().AddSubview(stackView)
		layout.PinEdgesToSuperView(stackView, foundation.EdgeInsets{Top: 10, Bottom: 10, Left: 20, Right: 20})

		w.Center()
		w.SetTitle("multi-roblox-macos")
		w.MakeKeyAndOrderFront(nil)

		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			clean()
			return true
		})
	})
}
