package app

import (
	"fmt"
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderActivityPanel builds the "Commit Activity" sparkline panel.
func renderActivityPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("📈 Commit Activity")

	// Fetch more commits for a better graph (use all available)
	dates, counts := groupCommitsByDate(m.Commits)

	if len(counts) == 0 {
		return theme.PanelStyle.Render(title + "\n  No commits found")
	}

	// Build sparkline from daily counts
	spark := renderSparkline(counts)

	// Show date labels below: first and last date
	firstDate := dates[0]
	lastDate := dates[len(dates)-1]
	dateRange := fmt.Sprintf("  %s %s %s",
		theme.SubtleStyle.Render(lastDate),
		strings.Repeat(" ", len(counts)),
		theme.SubtleStyle.Render(firstDate))

	body := fmt.Sprintf("  %s\n%s", spark, dateRange)
	return theme.PanelStyle.Render(title + "\n" + body)
}
