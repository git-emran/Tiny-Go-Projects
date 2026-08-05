package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/git-emran/tiny-go-project/tui-explorer/internal/theme"
)

func renderPane(p pane, width int, active bool, icons theme.IconSet) string {
	style := lipgloss.NewStyle().Width(width).Height(20).Padding(0, 1)

	if active {
		style = style.BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("212"))
	} else {
		style = style.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	}

	var b strings.Builder
	for i, e := range p.entries {
		icon := theme.IconFor(icons, e, i == p.cursor)
		text := icon + e.Name
		if e.IsDir {
			text += "/"
		}
		line := lipgloss.NewStyle().Foreground(theme.EntryColor(e.IsDir, false)).Render(text)
		if i == p.cursor {
			line = lipgloss.NewStyle().Reverse(true).Render(line)
		}
		b.WriteString(line + "\n")
	}
	return style.Render(b.String())
}

func renderTextPreview(content string, width int) string {
	style := lipgloss.NewStyle().Width(width).Height(20).Padding(0, 1).BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	return style.Render(content)
}

func (m Model) View() string {
	if m.width == 0 {
		return "loading"
	}
	colWidth := m.width / 3

	parentView := renderPane(m.parent, colWidth, false, m.icons)
	currentView := renderPane(m.current, colWidth, true, m.icons)

	var previewView string

	if m.previewText != "" {
		previewView = renderTextPreview(m.previewText, colWidth)
	} else {
		previewView = renderPane(m.preview, colWidth, false, m.icons)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, parentView, currentView, previewView)
}
