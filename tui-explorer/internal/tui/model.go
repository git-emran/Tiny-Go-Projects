// Package tui initializes and refresh panes

package tui

import (
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/git-emran/tiny-go-project/tui-explorer/internal/fs"
	"github.com/git-emran/tiny-go-project/tui-explorer/internal/theme"
)

type Model struct {
	cwd         string
	parent      pane
	current     pane
	preview     pane
	previewText string
	previewPath string
	width       int
	height      int
	pendingG    bool
	icons       theme.IconSet
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
		icons:   theme.IconSetnerd,
	}
	m, _ = m.refreshPanes()
	return m
}

func (m Model) SelectedPath() string {
	if len(m.current.entries) == 0 || m.current.cursor > len(m.current.entries) {
		return ""
	}
	sel := m.current.entries[m.current.cursor]
	return filepath.Join(m.cwd, sel.Name)
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

	m.preview = pane{}
	m.previewText = ""
	m.previewPath = ""

	if len(m.current.entries) == 0 || m.current.cursor >= len(m.current.entries) {
		return m, nil
	}

	sel := m.current.entries[m.current.cursor]
	selPath := filepath.Join(m.cwd, sel.Name)

	if sel.IsDir {
		m.preview = pane{path: selPath, entries: fs.ReadDir(selPath)}
		return m, nil
	}
	return m, loadPreviewCmd(selPath)
}
