package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type entry struct {
	name  string
	isDir bool
}

type model struct {
	cwd      string
	entries  []entry
	cursor   int
	pendingG bool
}

func readDir(path string) []entry {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var out []entry
	for _, f := range files {
		out = append(out, entry{name: f.Name(), isDir: f.IsDir()})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].isDir == out[j].isDir {
			return out[i].isDir
		}
		return strings.ToLower(out[i].name) < strings.ToLower(out[j].name)
	})

	return out
}

func initialModel() model {
	cwd, _ := os.Getwd()
	return model{
		cwd:     cwd,
		entries: readDir(cwd),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		// handle the pending "gg" sequence first
		if m.pendingG {
			m.pendingG = false
			if key == "g" {
				m.cursor = 0
				return m, nil
			}
			// fall through — not a real "gg", handle key normally below
		}

		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.cursor < len(m.entries)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "g":
			m.pendingG = true
		case "G":
			m.cursor = len(m.entries) - 1
		case "h", "left":
			m = m.goToParent()
		case "l", "right", "enter":
			m = m.enterSelected()
		case "ctrl+d":
			m.cursor = min(m.cursor+10, len(m.entries)-1)
		case "ctrl+u":
			m.cursor = max(m.cursor-10, 0)
		}
	}
	return m, nil
}

func (m model) goToParent() model {
	parent := filepath.Dir(m.cwd)
	if parent == m.cwd {
		return m
	}

	m.cwd = parent
	m.entries = readDir(parent)
	m.cursor = 0
	return m
}

func (m model) enterSelected() model {
	if len(m.entries) == 0 {
		return m
	}

	sel := m.entries[m.cursor]
	if sel.isDir {
		newPath := filepath.Join(m.cwd, sel.name)
		m.cwd = newPath
		m.entries = readDir(newPath)
		m.cursor = 0
	}
	return m
}

func (m model) View() string {
	var b strings.Builder
	b.WriteString(m.cwd + "\n\n")
	for i, e := range m.entries {
		cursor := " "
		if i == m.cursor {
			cursor = "> "
		}
		name := e.name
		if e.isDir {
			name += "/"
		}
		fmt.Fprintf(&b, "%s%s\n", cursor, name)
	}
	return b.String()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("error running program", err)
		os.Exit(1)
	}
}
