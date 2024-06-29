package open_app

import "testing"

func TestOpen(t *testing.T) {
	<-Open("/Applications/Roblox.app")
}
