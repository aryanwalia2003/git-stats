package theme

import "github.com/charmbracelet/lipgloss"

// Text styles for different stat elements
var (
	LabelStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cdd6f4")) // light gray/white text, removed bold for hierarchy

	ValueStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")). // clean white for readable numbers
		Bold(true)

	BarStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#a6e3a1")) // green specifically for bars/charts

	SubtleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#585b70")) // even dimmer text for secondary info
)
