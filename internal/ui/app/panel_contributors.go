package app

import (
	"fmt"
	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

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
		// Enforce fixed width for usernames to ensure charts align beautifully
		labels[i] = fmt.Sprintf("%-18s", m.Contributors[i].Label)
		values[i] = m.Contributors[i].Value
	}

	chart := renderBarChart(labels, values, 15) // max 15 chars wide
	return title + "\n\n" + chart + "\n"
}
