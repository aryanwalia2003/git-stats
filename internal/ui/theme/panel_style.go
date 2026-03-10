package theme

import "github.com/charmbracelet/lipgloss"

// Panel styles for the dashboard sections
var (
	PanelStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#45475a")). // Dimmer border to reduce dominance
		Padding(1, 2). // Consistent internal padding
		Width(48) // Fixed width for clean 2-column grid alignment

	FullWidthPanelStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#45475a")). // Dimmer border
		Padding(1, 2) // No width limit, takes available space

	PanelTitleStyle = lipgloss.NewStyle().
		Foreground(Primary).
		Bold(true).
		MarginBottom(1) // Clear separation from content
)
