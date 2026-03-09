package app

import tea "github.com/charmbracelet/bubbletea"

// Init fires when the program starts — kicks off the spinner + stat fetching.
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.Spinner.Tick,       // start spinner animation
		fetchStats(m.Reader), // fetch git stats in background
	)
}
