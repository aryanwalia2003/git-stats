package app

import (
	"fmt"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
	"github.com/charmbracelet/lipgloss"
)

// renderDashboard assembles the grid layout using layout primitives.
func (m Model) renderDashboard() string {
	head := GithubLogo() + "\n" + theme.TitleStyle.Render("📊 " + m.RepoName)
	r1 := flexRow(renderContributorsPanel(m), renderCommitsPanel(m))
	r2 := fullRow(renderActivityPanel(m))

	return lipgloss.NewStyle().Margin(1, 4).Render(
		fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s", head, r1, r2, m.renderBottomRows()),
	)
}

func (m Model) renderBottomRows() string {
	r3 := flexRow(renderChurnPanel(m), renderHourHeatmap(m))
	r4 := flexRow(renderVibesPanel(m), renderTimelinePanel(m))
	r5 := fullRow(renderDynamicsPanel(m))
	r6 := flexRow(renderDaysPanel(m), renderFilesPanel(m))

	foot := theme.SubtleStyle.Render("Press j/k or up/down to scroll. q to quit.")
	return fmt.Sprintf("%s\n\n%s\n\n%s\n\n%s\n\n%s\n", r3, r4, r5, r6, foot)
}
