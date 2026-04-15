package app

import (
	tea "charm.land/bubbletea/v2"
	"fmt"
)

type model struct {
	cursor   int
	choices  []string
	selected map[int]struct{}

	pageSize   int
	offset     int
	termHeight int
}

func InitialModel() model {
	packages, err := GetUpdates()
	if err != nil {
		packages =[]string{"Error: could not fetch updates"}
	}
	return model{
		choices:  packages,
		selected: make(map[int]struct{}),
		pageSize: 10, // fallback 
	}
}

func (m model) Init() tea.Cmd {
	return tea.RequestWindowSize
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.termHeight = msg.Height
		m.pageSize = max(1, m.termHeight-5)
		m.clampOffsetAndCursor()
		return m, nil

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
				m.adjustOffsetUp()
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
				m.adjustOffsetDown()
			}

		case "left", "h":
			if m.offset > 0 {
				m.offset = max(0, m.offset-m.pageSize)
				m.cursor = m.offset
			}

		case "right", "l":
			if m.offset+m.pageSize < len(m.choices) {
				m.offset = min(len(m.choices)-m.pageSize, m.offset+m.pageSize)
				m.cursor = m.offset
			}

		case "space":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "enter":
			selectedPackages := make([]string, 0, len(m.selected))
			for idx := range m.selected {
				selectedPackages = append(selectedPackages, m.choices[idx])
			}
			return m, tea.Sequence(
				tea.ClearScreen,
				RunUpdates(selectedPackages),
				tea.Quit,
			)
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	s := "What packages should we update? \n\n"

	end := min(m.offset+m.pageSize, len(m.choices))
	visible := m.choices[m.offset:end]

	for i, choice := range visible {
		absIdx := m.offset + i
		cursor := " "
		if m.cursor == absIdx {
			cursor = ">"
		}
		checked := " "
		if _, ok := m.selected[absIdx]; ok {
			checked = "x"
		}
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += fmt.Sprintf("\nPage %d/%d · ←/h prev · →/l next · ↑/↓ move · space select · enter confirm · q quit\n",
		(m.offset/m.pageSize)+1,
		(len(m.choices)+m.pageSize-1)/m.pageSize)

	v := tea.NewView(s)
	v.WindowTitle = "updarch"
	return v
}

func (m *model) adjustOffsetUp() {
	if m.cursor < m.offset {
		m.offset = m.cursor
	}
}

func (m *model) adjustOffsetDown() {
	if m.cursor >= m.offset+m.pageSize {
		m.offset = m.cursor - m.pageSize + 1
	}
}

func (m *model) clampOffsetAndCursor() {
	if m.offset < 0 {
		m.offset = 0
	}
	maxOffset := max(0, len(m.choices)-m.pageSize)
	if m.offset > maxOffset {
		m.offset = maxOffset
	}
	if m.cursor < 0 {
		m.cursor = 0
	}
	if m.cursor >= len(m.choices) {
		m.cursor = len(m.choices) - 1
	}
	if m.cursor < m.offset {
		m.offset = m.cursor
	} else if m.cursor >= m.offset+m.pageSize {
		m.offset = m.cursor - m.pageSize + 1
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

