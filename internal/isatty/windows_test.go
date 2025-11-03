//go:build windows

package isatty_test

import (
	"testing"

	"gh.tarampamp.am/colors/internal/isatty"
)

func TestCygwinPipeName(t *testing.T) {
	tests := []struct {
		name   string
		result bool
	}{
		{``, false},
		{`\msys-`, false},
		{`\cygwin-----`, false},
		{`\msys-x-PTY5-pty1-from-master`, false},
		{`\cygwin-x-PTY5-from-master`, false},
		{`\cygwin-x-pty2-from-toaster`, false},
		{`\cygwin--pty2-from-master`, false},
		{`\\cygwin-x-pty2-from-master`, false},
		{`\cygwin-x-pty2-from-master-`, true}, // for the feature
		{`\cygwin-e022582115c10879-pty4-from-master`, true},
		{`\msys-e022582115c10879-pty4-to-master`, true},
		{`\cygwin-e022582115c10879-pty4-to-master`, true},
		{`\Device\NamedPipe\cygwin-e022582115c10879-pty4-from-master`, true},
		{`\Device\NamedPipe\msys-e022582115c10879-pty4-to-master`, true},
		{`Device\NamedPipe\cygwin-e022582115c10879-pty4-to-master`, false},
	}

	for _, test := range tests {
		want, got := test.result, isatty.IsCygwinPipeName(test.name)

		if got != want {
			t.Fatalf("isatty(%q): got %v, want %v:", test.name, got, want)
		}
	}
}
