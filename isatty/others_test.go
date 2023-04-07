//go:build !windows
// +build !windows

package isatty_test

import (
	"os"
	"testing"

	"gh.tarampamp.am/colors/isatty"
)

func TestTerminal(t *testing.T) {
	// test for non-panic
	t.Log("os.Stdout:", isatty.IsTerminal(os.Stdout.Fd()))
}

func TestCygwinPipeName(t *testing.T) {
	if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		t.Fatal("should be false always")
	}
}
