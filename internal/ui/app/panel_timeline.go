package app

import (
	"fmt"
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderTimelinePanel builds a vertical timeline of major repo events.
func renderTimelinePanel(m Model) string {
	title := theme.PanelTitleStyle.Render("⏳  Project Timeline")

	events := computeTimeline(m.History)
	if len(events) == 0 {
		return title + "\n  No timeline events yet."
	}

	var sb strings.Builder
	sb.WriteString("\n")

	for i, ev := range events {
		// Event Header: "🌱 The Genesis (2024-03-04)"
		header := fmt.Sprintf("  %s %s %s",
			ev.Emoji,
			theme.ValueStyle.Render(ev.Title),
			theme.SubtleStyle.Render("("+ev.Date+")"))

		// Connection line unless it's the last item
		connector := "    │"
		if i == len(events)-1 {
			connector = ""
		}

		// Event Description
		desc := fmt.Sprintf("    %s", theme.LabelStyle.Render(ev.Description))

		sb.WriteString(header + "\n" + desc + "\n" + theme.SubtleStyle.Render(connector) + "\n")
	}

	return title + sb.String()
}
