package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/aryanwalia2003/git-stats/internal/git"
	"github.com/aryanwalia2003/git-stats/internal/ui/app"
)

func main() {
	// Get the current working directory (the repo we're inside)
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Create a git reader pointing at the current directory
	reader := git.NewReader(cwd)

	// Read repo info from .git/config
	repo, err := reader.GetCurrentRepo()
	if err != nil {
		fmt.Println("Not a git repo or no remote set:", err)
		os.Exit(1)
	}

	// Start the TUI with the real repo name and alternate screen for scrolling
	initialModel := app.New(repo.Name, reader)
	program := tea.NewProgram(initialModel, tea.WithAltScreen())

	if _, err := program.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
		os.Exit(1)
	}
}

