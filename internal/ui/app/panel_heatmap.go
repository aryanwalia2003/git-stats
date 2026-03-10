package app

import (
	"fmt"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderHourHeatmap shows commit frequency across 24 hours as a sparkline.
func renderHourHeatmap(m Model) string {
	title := theme.PanelTitleStyle.Render("🕐 Hour-of-Day Heatmap")

	histogram := computeHourHistogram(m.Commits)
	spark := renderSparkline(histogram)

	// Hour labels: 0  3  6  9  12 15 18 21
	labels := "  " + theme.SubtleStyle.Render("0  3  6  9  12 15 18 21")

	return title + "\n  " + spark + "\n" + labels
}

// renderDayOfWeekChart shows commits per weekday as a bar chart.
func renderDayOfWeekChart(m Model) string {
	title := theme.PanelTitleStyle.Render("📅 Day-of-Week")

	histogram := computeDayHistogram(m.Commits)
	dayLabels := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	chart := ""
	for i, label := range dayLabels {
		chart += fmt.Sprintf("  %s %s\n",
			theme.LabelStyle.Render(fmt.Sprintf("%-3s", label)),
			renderSparkline([]int{histogram[i]}))
	}

	return title + "\n" + chart
}
