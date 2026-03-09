package app

import (
	"fmt"
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/domain"
	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// formatCommitDrilldown creates a 2-line drilldown view for a commit.
// Returns: "a1b2c3d (42 lines)\n  └─ added login (2024-03-04)"
func formatCommitDrilldown(c domain.Stat, count int, suffix string) string {
	if c.RepoID == "" {
		return theme.SubtleStyle.Render("N/A")
	}

	// Shorten message if too long, and remove any stray newlines
	msg := sanitizeMessage(c.Message)
	if len(msg) > 25 {
		msg = msg[:22] + "..."
	}

	// Extract just the date (YYYY-MM-DD)
	date := c.Date
	if len(date) > 10 {
		date = date[:10]
	}

	line1 := fmt.Sprintf("%s (%d %s)", theme.ValueStyle.Render(c.RepoID), count, suffix)
	line2 := theme.SubtleStyle.Render(fmt.Sprintf("    └─ %s (%s)", msg, date))

	return line1 + "\n" + line2
}

// sanitizeMessage removes newlines and carriage returns from commit subjects.
func sanitizeMessage(msg string) string {
	return strings.TrimSpace(strings.ReplaceAll(msg, "\n", " "))
}
