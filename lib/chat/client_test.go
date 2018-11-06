package chat

import (
	"testing"
)

func TestNewHandler(t *testing.T) {
	handler := newHandler("TeHpEnGuInOfDoOm")
	if handler.channel != "tehpenguinofdoom" {
		t.Errorf("Expected (tehpenguinofdoom), got (%s)", handler.channel)
	}
}
