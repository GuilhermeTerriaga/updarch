package app

import (
	"testing"
	tea "charm.land/bubbletea/v2"
)

func TestModelUpdateWindowSize(t *testing.T) {
	m := InitialModel()
	// Simulate a window size message
	newModel, cmd := m.Update(tea.WindowSizeMsg{Height: 30})
	m2 := newModel.(model)

	if m2.termHeight != 30 {
		t.Errorf("termHeight = %d, want 30", m2.termHeight)
	}
	expectedPageSize := 30 - 5 // 25
	if m2.pageSize != expectedPageSize {
		t.Errorf("pageSize = %d, want %d", m2.pageSize, expectedPageSize)
	}
	if cmd != nil {
		t.Errorf("expected nil cmd, got %v", cmd)
	}
}

