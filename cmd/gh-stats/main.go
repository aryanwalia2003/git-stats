package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/yourusername/gh-stats/internal/ui/app"
)

func main() {
	initialModel := app.New("Initializing...") //model bana rhe naya with name initializing
	program := tea.NewProgram(initialModel)

	if _, err := program.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}
