package app

import (
	"fmt"
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderDynamicsPanel builds the "Team Dynamics & Branches" panel.
func renderDynamicsPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("👥  Team Dynamics & Branches")

	var body strings.Builder
	body.WriteString(fmt.Sprintf("  🚢 Biggest Merge:\n  %s\n\n", computeBiggestMerge(m.Merges)))

	lifecycles := computeLifecycles(m.History)
	if len(lifecycles) == 0 {
		body.WriteString(theme.SubtleStyle.Render("  No contributor data found."))
		return title + "\n" + body.String()
	}

	for _, lc := range lifecycles {
		name := theme.LabelStyle.Render(lc.Name)
		joined := theme.SubtleStyle.Render(fmt.Sprintf("Joined %s", lc.JoinDate))
		
		status := theme.BarStyle.Render("🟢 Active")
		if lc.LeftDate != "" {
			status = theme.SubtleStyle.Render(fmt.Sprintf("💤 Sabbatical since %s", lc.LeftDate))
		}
		
		// 1 space for padding, name, status. The L-bracket aligns under the name.
		body.WriteString(fmt.Sprintf(" %s %s\n   └─ %s\n", name, status, joined))
	}

	return title + "\n" + body.String()
}
