package app

import (
	"testing"
)

func TestToSlice(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
	}{
		{
			name:  "single package",
			input: "firefox 123.0-1\n",
			want:  []string{"firefox"},
		},
		{
			name:  "multiple packages",
			input: "firefox 123.0-1\nchromium 124.0-1\n",
			want:  []string{"firefox", "chromium"},
		},
		{
			name:  "ignore lines without package name",
			input: "\n\nfirefox 123.0-1\n\n",
			want:  []string{"firefox"},
		},
		{
			name:  "empty input",
			input: "",
			want:  []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := toSlice(tt.input)
			if len(got) != len(tt.want) {
				t.Errorf("toSlice() length = %d, want %d", len(got), len(tt.want))
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("toSlice()[%d] = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}

func TestRunUpdates(t *testing.T) {
	// RunUpdates returns a tea.Cmd that we cannot easily inspect.
	// We just verify it does not panic and returns something non-nil.
	packages := []string{"firefox", "chromium"}
	cmd := RunUpdates(packages)
	if cmd == nil {
		t.Error("RunUpdates returned nil tea.Cmd")
	}
}

