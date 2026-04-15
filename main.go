package main

import (
	"fmt"
	"os"
	"updarch/app"

	tea "charm.land/bubbletea/v2"
)

func main() {
	p := tea.NewProgram(app.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}
}
