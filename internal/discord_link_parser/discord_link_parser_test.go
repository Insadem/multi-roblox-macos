package discord_link_parser

import (
	"fmt"
	"testing"
)

func TestGetDiscordLink(t *testing.T) {
	link := DiscordLink()
	if link == "" {
		t.Errorf("discord link is empty")
	}

	fmt.Println(link)
}
