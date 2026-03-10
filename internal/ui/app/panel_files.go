package app

import (
	"fmt"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderFilesPanel builds the "Most/Least Touched Files" layout.
func renderFilesPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("📁 File Hotspots")

	if len(m.Files) == 0 {
		return title + "\n  No files tracked"
	}

	// Get top 3 most edited files
	limit := 3
	if len(m.Files) < limit {
		limit = len(m.Files)
	}

	labels, vals := make([]string, limit), make([]int, limit)
	for i := 0; i < limit; i++ {
		labels[i] = m.Files[i].Label
		vals[i] = m.Files[i].Value
	}

	least := ""
	for i := len(m.Files) - 1; i >= 0 && i >= len(m.Files)-limit; i-- {
		lbl := m.Files[i].Label
		if len(lbl) > 20 { lbl = lbl[:17] + "..."}
		least += fmt.Sprintf("  %s %s\n", theme.SubtleStyle.Render("1 edit"), theme.SubtleStyle.Render(lbl))
	}

	return fmt.Sprintf("%s\n%s\n%s\n\n%s\n%s", title,
		theme.ValueStyle.Render("  🔥 Heaviest Modification"), renderBarChart(labels, vals, 15),
		theme.ValueStyle.Render("  🧊 Barely Touched"), least)
}
