package app

import "github.com/aryanwalia2003/git-stats/internal/ui/theme"

// renderContributorsPanel builds the "Top Contributors" panel with bar chart.
func renderContributorsPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("👥 Top Contributors")

	limit := 5
	if len(m.Contributors) < limit {
		limit = len(m.Contributors)
	}

	labels := make([]string, limit)
	values := make([]int, limit)
	for i := 0; i < limit; i++ {
		labels[i] = m.Contributors[i].Label
		values[i] = m.Contributors[i].Value
	}

	chart := renderBarChart(labels, values, 20) // max 20 chars wide
	return theme.PanelStyle.Render(title + "\n" + chart)
}
