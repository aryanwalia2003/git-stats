package app

import (
	"fmt"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderFilesPanel builds the "Most/Least Touched Files" layout.
func renderFilesPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("📁 File Hotspots")

	if len(m.Files) == 0 {
		return theme.PanelStyle.Render(title + "\n  No files tracked")
	}

	// Get top 3 most edited files
	most := ""
	limit := 3
	if len(m.Files) < limit { limit = len(m.Files) }
	
	for i := 0; i < limit; i++ {
		f := m.Files[i]
		most += fmt.Sprintf("  %s %s\n", 
			theme.LabelStyle.Render(fmt.Sprintf("%d edits ", f.Value)), 
			theme.SubtleStyle.Render(f.Label))
	}

	// Get top 3 least edited files (from the bottom of the sorted slice)
	least := ""
	for i := len(m.Files) - 1; i >= 0 && i >= len(m.Files)-limit; i-- {
		f := m.Files[i]
		least += fmt.Sprintf("  %s %s\n", 
			theme.LabelStyle.Render(fmt.Sprintf("%d edits ", f.Value)), 
			theme.SubtleStyle.Render(f.Label))
	}

	body := fmt.Sprintf("\n%s\n%s\n\n%s\n%s", 
		theme.ValueStyle.Render("  🔥 Heaviest Modification"), most,
		theme.ValueStyle.Render("  🧊 Barely Touched"), least)

	return theme.PanelStyle.Render(title + body)
}
