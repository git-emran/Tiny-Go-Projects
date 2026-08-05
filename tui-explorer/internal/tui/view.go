package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/git-emran/tiny-go-project/tui-explorer/internal/theme"
)

// truncate clips s to at most maxRunes visible runes (ignores ANSI, so call
// this on plain text BEFORE applying lipgloss colour styles).
func truncate(s string, maxRunes int) string {
	if maxRunes <= 0 {
		return ""
	}
	r := []rune(s)
	if len(r) > maxRunes {
		return string(r[:maxRunes])
	}
	return s
}

func renderPane(p pane, width int, height int, active bool, icons theme.IconSet) string {
	// visible lines = terminal height minus top+bottom border
	visible := height - 2
	if visible < 1 {
		visible = 1
	}
	// inner text width = pane width minus left+right border (2) minus left+right padding (2)
	inner := width - 4
	if inner < 1 {
		inner = 1
	}

	style := lipgloss.NewStyle().
		Width(width).Height(visible).
		MaxWidth(width+2).MaxHeight(visible+2). // hard ceiling incl. border
		Padding(0, 1)

	if active {
		style = style.BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("212"))
	} else {
		style = style.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	}

	// Compute scroll offset so the cursor is always visible
	offset := p.offset
	if p.cursor < offset {
		offset = p.cursor
	}
	if p.cursor >= offset+visible {
		offset = p.cursor - visible + 1
	}

	var b strings.Builder
	for i := offset; i < len(p.entries) && i < offset+visible; i++ {
		e := p.entries[i]
		icon := theme.IconFor(icons, e, i == p.cursor)
		name := e.Name
		if e.IsDir {
			name += "/"
		}
		// Truncate plain text to inner width before colouring
		text := truncate(icon+name, inner)
		line := lipgloss.NewStyle().Foreground(theme.EntryColor(e.IsDir, false)).Render(text)
		if i == p.cursor {
			line = lipgloss.NewStyle().Reverse(true).Render(text)
		}
		b.WriteString(line + "\n")
	}
	return style.Render(b.String())
}

func renderTextPreview(content string, width int, height int) string {
	visible := height - 2
	if visible < 1 {
		visible = 1
	}
	inner := width - 4
	if inner < 1 {
		inner = 1
	}

	lines := strings.Split(content, "\n")
	var b strings.Builder
	count := 0
	for _, l := range lines {
		if count >= visible {
			break
		}
		b.WriteString(truncate(l, inner) + "\n")
		count++
	}

	style := lipgloss.NewStyle().
		Width(width).Height(visible).
		MaxWidth(width+2).MaxHeight(visible+2). // hard ceiling incl. border
		Padding(0, 1).
		BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	return style.Render(b.String())
}

func (m Model) View() string {
	if m.width == 0 {
		return "loading"
	}
	colWidth := m.width / 3

	parentView := renderPane(m.parent, colWidth, m.height, false, m.icons)
	currentView := renderPane(m.current, colWidth, m.height, true, m.icons)

	var previewView string
	if m.previewText != "" {
		previewView = renderTextPreview(m.previewText, colWidth, m.height)
	} else {
		previewView = renderPane(m.preview, colWidth, m.height, false, m.icons)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, parentView, currentView, previewView)
}
