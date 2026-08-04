// Package tui initializes and refresh panes

package tui

import (
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/git-emran/tiny-go-project/tui-explorer/internal/fs"
)

type Model struct {
	cwd      string
	parent   pane
	current  pane
	preview  pane
	width    int
	height   int
	pendingG bool
}

type pane struct {
	path    string
	entries []fs.Entry
	cursor  int
}

func InitialModel() Model {
	cwd, _ := os.Getwd()
	m := Model{
		cwd:     cwd,
		current: pane{path: cwd, entries: fs.ReadDir(cwd)},
	}
	m, _ = m.refreshPanes()
	return m
}

func (m Model) refreshPanes() (Model, tea.Cmd) {
	m.parent = pane{
		path:    filepath.Dir(m.cwd),
		entries: fs.ReadDir(filepath.Dir(m.cwd)),
	}
	for i, e := range m.parent.entries {
		if e.Name == filepath.Base(m.cwd) {
			m.parent.cursor = i
		}
	}

	if len(m.current.entries) > 0 {
		sel := m.current.entries[m.current.cursor]
		selPath := filepath.Join(m.cwd, sel.Name)

		if sel.IsDir {
			m.preview = pane{path: selPath, entries: fs.ReadDir(selPath)}
		} else {
			m.preview = pane{}
		}
	} else {
		m.preview = pane{}
	}
	return m, nil
}
