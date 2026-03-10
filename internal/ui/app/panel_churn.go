package app

import (
	"fmt"
	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderChurnPanel builds the "Code Churn" summary panel.
func renderChurnPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("🌪️  Code Churn")

	stats := computeChurnStats(m.Churn)

	body := fmt.Sprintf("  %s %s   %s %s   %s %s\n",
		theme.ValueStyle.Render(fmt.Sprintf("+%d", stats.TotalAdded)),
		theme.SubtleStyle.Render("added"),
		theme.ValueStyle.Render(fmt.Sprintf("-%d", stats.TotalDeleted)),
		theme.SubtleStyle.Render("deleted"),
		theme.LabelStyle.Render(fmt.Sprintf("%d", len(m.Churn))),
		theme.SubtleStyle.Render("commits"))

	body += fmt.Sprintf("  %s\n\n", renderRatioBar(stats.TotalAdded, stats.TotalDeleted, 40))

	body += fmt.Sprintf("\n %s\n %s\n\n", 
		theme.ValueStyle.Render(fmt.Sprintf("%d touched", stats.TotalFiles)) + theme.SubtleStyle.Render(" files"),
		theme.ValueStyle.Render(fmt.Sprintf("%d lines, %d files", stats.AvgLines, stats.AvgFiles)) + theme.SubtleStyle.Render(" per commit average"))

	body += fmt.Sprintf(" %s\n %s\n\n", theme.LabelStyle.Render("🚀 Biggest Commit:"), formatCommitDrilldown(stats.BiggestCommit, stats.BiggestCommit.Value+stats.BiggestCommit.Value2, "lines"))
	body += fmt.Sprintf(" %s\n %s\n\n", theme.LabelStyle.Render("📁 Most Files:"), formatCommitDrilldown(stats.BiggestCommitFiles, stats.BiggestCommitFiles.Value3, "files"))
	body += fmt.Sprintf(" %s\n %s\n\n", theme.LabelStyle.Render("🐭 Smallest Commit:"), formatCommitDrilldown(stats.SmallestCommit, stats.SmallestCommit.Value+stats.SmallestCommit.Value2, "lines"))
	body += fmt.Sprintf(" %s\n %s\n", theme.LabelStyle.Render("🔧 Big Refactor:"), formatCommitDrilldown(stats.BiggestRefactor, stats.BiggestRefactor.Value2, "lines deleted"))

	return title + "\n" + body
}
