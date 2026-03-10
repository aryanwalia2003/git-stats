package app

import (
	"fmt"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderCommitsPanel builds the "Recent Commits" panel.
func renderCommitsPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("📝 Recent Commits")
	body := ""

	limit := 5
	if len(m.Commits) < limit {
		limit = len(m.Commits)
	}

	for i := 0; i < limit; i++ {
		c := m.Commits[i]
		msg := sanitizeMessage(c.Message)
		if len(msg) > 30 {
			msg = msg[:27] + "..." // Truncate so it fits the panel
		}

		// Metadata is dim, Message is bright/primary content
		line := fmt.Sprintf(" %s %s %s\n   └─ %s",
			theme.SubtleStyle.Render(c.RepoID),    // short hash
			theme.SubtleStyle.Render(c.Date[:10]), // date first
			theme.SubtleStyle.Render(c.Label),     // author name dimmed
			theme.ValueStyle.Render(msg))          // message pops
		body += line + "\n"
	}

	return title + "\n\n" + body
}
