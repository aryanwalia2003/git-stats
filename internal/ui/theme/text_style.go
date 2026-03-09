package theme

import "github.com/charmbracelet/lipgloss"

// Text styles for different stat elements
var (
	LabelStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cdd6f4")). // light gray text
		Bold(true)

	ValueStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a6e3a1")). // green for numbers
		Bold(true)

	SubtleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6c7086")) // dimmed text for secondary info
)
