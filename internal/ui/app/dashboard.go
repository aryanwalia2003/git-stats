package app

import (
	"fmt"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
	"github.com/charmbracelet/lipgloss"
)

// renderDashboard assembles the 3-row grid layout of the dashboard.
func (m Model) renderDashboard() string {
	header := theme.TitleStyle.Render("📊 " + m.RepoName)

	row1 := lipgloss.JoinHorizontal(lipgloss.Top,
		renderContributorsPanel(m), "  ", renderCommitsPanel(m))

	// Activity gets full width for the 6-month calendar
	row2 := renderActivityPanel(m)

	row3 := lipgloss.JoinHorizontal(lipgloss.Top,
		renderChurnPanel(m), "  ", renderHourHeatmap(m))

	row4 := lipgloss.JoinHorizontal(lipgloss.Top,
		renderVibesPanel(m), "  ", renderTimelinePanel(m))

	row5 := renderDynamicsPanel(m)

	row6 := lipgloss.JoinHorizontal(lipgloss.Top,
		renderDaysPanel(m), "  ", renderFilesPanel(m))

	footer := theme.SubtleStyle.Render("Press j/k or up/down to scroll. q to quit.")

	content := fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n\n%s\n",
		header, row1, row2, row3, row4, row5, row6, footer)

	// Add margins all around the dashboard for breathing room
	return lipgloss.NewStyle().Margin(1, 4).Render(content)
}
