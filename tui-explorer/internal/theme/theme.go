// Package theme provides color scheme
package theme

import "github.com/charmbracelet/lipgloss"

var (
	Primary   = lipgloss.Color("212") // active pane border selection
	Muted     = lipgloss.Color("240") // inactive border
	DirColor  = lipgloss.Color("39")
	FileColor = lipgloss.Color("252") // Directories
	ExecColor = lipgloss.Color("42")  // Executables
	Cursor    = lipgloss.NewStyle().Reverse(true)

	ActivePane = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(Primary)

	InactivePane = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(Muted)

	StatusBar = lipgloss.NewStyle().Background(lipgloss.Color("236")).Foreground(lipgloss.Color("252")).Padding(0, 1)
)

func EntryColor(isDir, isExec bool) lipgloss.Color {
	switch {
	case isDir:
		return DirColor
	case isExec:
		return ExecColor
	default:
		return FileColor
	}
}
