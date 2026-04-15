package app

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	tea "charm.land/bubbletea/v2"
)

var execCommand = exec.Command

func GetUpdates() ([]string, error) {
	updates, err := execCommand("paru", "-Qu").Output()
	if err != nil {
		return nil, err
	}
	output := string(updates)
	packages := toSlice(output)
	return packages, nil
}

func toSlice(txt string) []string {
	var packs []string
	scan := bufio.NewScanner(strings.NewReader(txt))
	for scan.Scan() {
		parts := strings.Fields(scan.Text())
		if len(parts) > 0 {
			packs = append(packs, parts[0])
		}
	}
	fmt.Print(packs)
	return packs
}

func RunUpdates(packages []string) tea.Cmd {
	args := append([]string{"-S"}, packages...)
	cmd := execCommand("paru", args...)
	return tea.ExecProcess(cmd, nil)
}
