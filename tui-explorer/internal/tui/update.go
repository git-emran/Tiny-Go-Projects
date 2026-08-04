package tui

import (
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/git-emran/tiny-go-project/tui-explorer/internal/fs"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		key := msg.String()
		var cmd tea.Cmd

		// handle the pending "gg" sequence first
		if m.pendingG {
			m.pendingG = false
			if key == "g" {
				m.current.cursor = 0
				m, cmd = m.refreshPanes()
				return m, cmd
			}
			// fall through — not a real "gg", handle key normally below
		}

		switch key {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.current.cursor < len(m.current.entries)-1 {
				m.current.cursor++
				m, cmd = m.refreshPanes()
			}
		case "k", "up":
			if m.current.cursor > 0 {
				m.current.cursor--
				m, cmd = m.refreshPanes()
			}
		case "g":
			m.pendingG = true
		case "G":
			m.current.cursor = len(m.current.entries) - 1
			m, cmd = m.refreshPanes()
		case "h", "left":
			m, cmd = m.goToParent()
		case "l", "right", "enter":
			m, cmd = m.enterSelected()
		case "ctrl+d":
			m.current.cursor = min(m.current.cursor+10, len(m.current.entries)-1)
			m, cmd = m.refreshPanes()
		case "ctrl+u":
			m.current.cursor = max(m.current.cursor-10, 0)
			m, cmd = m.refreshPanes()

		}
	}
	return m, nil
}

func (m Model) goToParent() (Model, tea.Cmd) {
	parent := filepath.Dir(m.cwd)
	if parent == m.cwd {
		return m, nil
	}

	m.cwd = parent
	m.current = pane{path: parent, entries: fs.ReadDir(parent)}
	return m.refreshPanes()
}

func (m Model) enterSelected() (Model, tea.Cmd) {
	if len(m.current.entries) == 0 {
		return m, nil
	}

	sel := m.current.entries[m.current.cursor]
	if sel.IsDir {
		newPath := filepath.Join(m.cwd, sel.Name)
		m.cwd = newPath
		m.current = pane{path: newPath, entries: fs.ReadDir(newPath)}
		m.current.cursor = 0
		return m.refreshPanes()
	}
	return m, nil
}
