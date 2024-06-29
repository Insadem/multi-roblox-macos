# multi_roblox_macos

#### Info
Roblox Multi Instance now on MacOs! \
This is part of my [project](https://github.com/users/Insadem/projects/2) that strides bring Roblox customization to Mac. \
You can grab latest release [here](), no need to build yourself unless you have intel/amd mac.

#### Hardware Support
m1/m2 (any silicone ones) - 100% compatibility. \
intel/amd - should work too, though can't compile myself 'cus I haven't intel/amd mac.

#### How To Build Yourself
Set up [fyne](https://docs.fyne.io/started/) first. \ 
Then run this in terminal: \
`go install fyne.io/fyne/v2/cmd/fyne@latest` \
`fyne get https://github.com/Insadem/multi_roblox_macos`

If you build from source (this repo): \
`fyne package -os darwin -icon ./resources/app_icon.png`