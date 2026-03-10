package app

import (
	"fmt"
	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderChurnPanel builds the "Code Churn" summary panel.
func renderChurnPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("🌪️  Code Churn")

	stats := computeChurnStats(m.Churn)

	body := fmt.Sprintf("  %s %s   %s %s   %s %s\n\n",
		theme.ValueStyle.Render(fmt.Sprintf("+%d", stats.TotalAdded)),
		theme.SubtleStyle.Render("added"),
		theme.ValueStyle.Render(fmt.Sprintf("-%d", stats.TotalDeleted)),
		theme.SubtleStyle.Render("deleted"),
		theme.LabelStyle.Render(fmt.Sprintf("%d", len(m.Churn))),
		theme.SubtleStyle.Render("commits"))

	body += fmt.Sprintf("  📂 Total Files:    %s\n", theme.ValueStyle.Render(fmt.Sprintf("%d touched", stats.TotalFiles)))
	body += fmt.Sprintf("  📏 Avg/Commit:     %s\n\n", theme.ValueStyle.Render(fmt.Sprintf("%d lines, %d files", stats.AvgLines, stats.AvgFiles)))

	body += fmt.Sprintf("  🚀 Biggest Commit:\n  %s\n\n", formatCommitDrilldown(stats.BiggestCommit, stats.BiggestCommit.Value+stats.BiggestCommit.Value2, "lines"))
	body += fmt.Sprintf("  📁 Most Files:\n  %s\n\n", formatCommitDrilldown(stats.BiggestCommitFiles, stats.BiggestCommitFiles.Value3, "files"))
	body += fmt.Sprintf("  🐭 Smallest Commit:\n  %s\n\n", formatCommitDrilldown(stats.SmallestCommit, stats.SmallestCommit.Value+stats.SmallestCommit.Value2, "lines"))
	body += fmt.Sprintf("  🔧 Big Refactor:\n  %s\n", formatCommitDrilldown(stats.BiggestRefactor, stats.BiggestRefactor.Value2, "lines deleted"))

	return title + "\n" + body
}
