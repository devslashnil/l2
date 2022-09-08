package main

import (
	"testing"
	"time"
)

// TestDev01 проверяет формат и часовой пояс GetTime()
func TestDev01(t *testing.T) {
	got := GetTime()
	want := time.Now().UTC().Format(time.Kitchen)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
