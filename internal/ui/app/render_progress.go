package app

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// renderProgressBar draws a percentage bar like [██████····]
func renderProgressBar(percentage float64, width int, color lipgloss.Color) string {
	filled := int((percentage / 100.0) * float64(width))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}
	empty := width - filled

	style := lipgloss.NewStyle().Foreground(color)
	subtle := lipgloss.NewStyle().Foreground(lipgloss.Color("241")) // Gray

	bar := style.Render(strings.Repeat("█", filled)) + subtle.Render(strings.Repeat("·", empty))
	return fmt.Sprintf("[%s]", bar)
}

// renderRatioBar draws a balanced + vs - bar.
func renderRatioBar(added, deleted, width int) string {
	total := added + deleted
	if total == 0 {
		return lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render(strings.Repeat("·", width))
	}

	addWidth := int((float64(added) / float64(total)) * float64(width))
	delWidth := width - addWidth // fill the rest

	addStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("42"))  // Green
	delStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("196")) // Red

	return addStyle.Render(strings.Repeat("█", addWidth)) + delStyle.Render(strings.Repeat("█", delWidth))
}
