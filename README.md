# multi_roblox_macos

#### Info
Roblox Multi Instance is now on MacOs! \
This is part of my [project](https://github.com/users/Insadem/projects/2) that strides bring Roblox customization to Mac. \
Tutorial about how to run is here []()

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