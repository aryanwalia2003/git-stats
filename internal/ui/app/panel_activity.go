package app

import (
	"fmt"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
	"github.com/charmbracelet/lipgloss"
)

// renderActivityPanel builds the "Commit Activity" sparkline panel.
func renderActivityPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("📈 Commit Activity")

	// We want to show exactly the last 6 months in distinct blocks
	blocks := computeMonthlyBlocks(m.History, 6)

	if len(blocks) == 0 {
		return title + "\n  No commits found"
	}

	calendar := renderMonthBlocks(blocks)

	legend := fmt.Sprintf("    Less %s %s %s %s %s More",
		lipgloss.NewStyle().Foreground(theme.CalendarColors[0]).Render("■"),
		lipgloss.NewStyle().Foreground(theme.CalendarColors[1]).Render("■"),
		lipgloss.NewStyle().Foreground(theme.CalendarColors[2]).Render("■"),
		lipgloss.NewStyle().Foreground(theme.CalendarColors[3]).Render("■"),
		lipgloss.NewStyle().Foreground(theme.CalendarColors[4]).Render("■"))

	body := fmt.Sprintf("\n%s\n%s", calendar, legend)
	return title + body
}
