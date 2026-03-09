package app

import (
	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderDaysPanel builds the "Day of Week" bar chart panel.
func renderDaysPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("📅 Day Distribution")

	if len(m.History) == 0 {
		return theme.PanelStyle.Render(title + "\n  No commits found")
	}

	// We calculate days across the entire parsed history
	days := make([]int, 7) // Sun=0 through Sat=6
	for _, c := range m.History {
		day := extractWeekday(c.Date)
		days[day]++
	}

	labels := []string{
		"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday",
	}

	// renderBarChart takes labels, values, and a max width for the bars
	chart := renderBarChart(labels, days, 20)

	body := "\n" + chart
	return theme.PanelStyle.Render(title + body)
}
