# multi_roblox_macos

#### Info
Roblox Multi Instance is now on MacOs! \
This is part of my [project](https://github.com/users/Insadem/projects/2) that strides bring Roblox customization to Mac. \
Video tutorial is [here](https://www.youtube.com/watch?v=2M1Gk0WrM2w).

#### Hardware Support
m1/m2 (any silicone ones) - 100% compatibility. \
intel/amd - should work too, though didn't test.

#### How To Build Yourself
Set up [fyne](https://docs.fyne.io/started/) first. 
Then run in the terminal: \
`go install fyne.io/fyne/v2/cmd/fyne@latest` \
`fyne get https://github.com/Insadem/multi_roblox_macos`

If you build from source (this repo): \
`fyne package -os darwin -icon ./resources/app_icon.png`

To build for amd/intel mac from silicone one (unix file, pack to app yourself): \
`GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build`
