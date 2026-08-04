package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/git-emran/tiny-go-project/tui-explorer/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("error running program", err)
		os.Exit(1)
	}
}
