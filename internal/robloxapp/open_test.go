package robloxapp

import (
	"testing"
	"time"
)

func TestOpen(t *testing.T) {
	clear, err := Open()
	if err != nil {
		t.Error(t)
		return
	}

	defer clear()
	<-time.After(time.Second * 2)
}
