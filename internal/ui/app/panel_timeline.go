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
		// Emojis vary in width, so we format them + pad to ensure stable alignment
		emojiSpace := fmt.Sprintf("%s ", ev.Emoji) 
		header := fmt.Sprintf(" %s %s %s",
			emojiSpace,
			theme.LabelStyle.Render(ev.Title),     // Neutral title
			theme.SubtleStyle.Render("("+ev.Date+")")) // Dim date

		// Connection line unless it's the last item
		connector := "   │"
		if i == len(events)-1 {
			connector = ""
		}

		// Event Description
		desc := fmt.Sprintf("   %s", theme.ValueStyle.Render(ev.Description))

		sb.WriteString(header + "\n" + desc + "\n" + theme.SubtleStyle.Render(connector) + "\n")
	}

	return title + sb.String()
}
