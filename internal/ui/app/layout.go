package app

import (
	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
	"github.com/charmbracelet/lipgloss"
)

// flexRow wraps two strings in equal-height borders and joins them horizontally.
func flexRow(left, right string) string {
	leftHeight := lipgloss.Height(left)
	rightHeight := lipgloss.Height(right)

	max := leftHeight
	if rightHeight > max {
		max = rightHeight
	}

	leftPanel := theme.PanelStyle.Render(lipgloss.NewStyle().Height(max).Render(left))
	rightPanel := theme.PanelStyle.Render(lipgloss.NewStyle().Height(max).Render(right))

	return lipgloss.JoinHorizontal(lipgloss.Top, leftPanel, "  ", rightPanel)
}

// fullRow wraps content in a full-width border.
func fullRow(content string) string {
	return theme.FullWidthPanelStyle.Render(content)
}
