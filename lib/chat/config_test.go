package chat

import (
	"testing"
)

func TestBotHandler(t *testing.T) {
	bot := botConfig{"UserName", "password", "ChannelName"}
	handler := bot.Handler()
	if handler.channel != "channelname" {
		t.Errorf("Expected (channelname), got (%s)", handler.channel)
	}
	return
}
