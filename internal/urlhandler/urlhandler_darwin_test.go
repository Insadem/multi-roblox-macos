package urlhandler

import "testing"

func TestUrlHandler(t *testing.T) {
	h, err := New()
	if err != nil {
		t.Error(err)
		return
	}

	v := h.Check(ROBLOX_BUNDLE_IDENTIFIER, "roblox")
	if v != true {
		t.Error("roblox handler supposed to be set to default")
	}

	v = h.Check(ROBLOX_BUNDLE_IDENTIFIER, "roblox-player")
	if v != true {
		t.Error("roblox-player handler supposed to be set to default")
	}

	ok := h.Set("com.Roblox.RobloxStudio", "roblox-player")
	if !ok {
		t.Error("set expected to be ok")
	}

	v = h.Check(ROBLOX_BUNDLE_IDENTIFIER, "roblox-player")
	if v == true {
		t.Error("roblox-player handler supposed to be not set")
	}

	h.Set(ROBLOX_BUNDLE_IDENTIFIER, "roblox-player")
	v = h.Check(ROBLOX_BUNDLE_IDENTIFIER, "roblox-player")
	if v != true {
		t.Error("roblox-player handler supposed to be set to default")
	}

	v = h.Check(ROBLOX_BUNDLE_IDENTIFIER, "non-existant")
	if v == true {
		t.Error("supposed to be not set")
	}
}
