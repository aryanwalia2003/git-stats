package app

import (
	"fmt"
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderBarChart draws a horizontal bar chart for the given labels and values.
// Example output:
//   aryanwalia2003  ████████████████ 42
//   contributor2    ████████         21
func renderBarChart(labels []string, values []int, maxWidth int) string {
	if len(labels) == 0 {
		return ""
	}

	// Find max value for scaling
	maxVal := 0
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	if maxVal == 0 {
		maxVal = 1
	}

	var lines []string
	for i, label := range labels {
		// Scale the bar width relative to max
		barLen := (values[i] * maxWidth) / maxVal
		if barLen == 0 && values[i] > 0 {
			barLen = 1 // at least 1 block for non-zero values
		}
		
		// Use BarStyle for chart bars instead of ValueStyle
		bar := theme.BarStyle.Render(strings.Repeat("█", barLen))
		count := theme.ValueStyle.Render(fmt.Sprintf("%d", values[i]))
		
		// Ensure labels are strictly 18 chars wide for perfect alignment of bars and numbers
		name := theme.LabelStyle.Render(fmt.Sprintf("%-18s", label))
		
		// Add subtle leading whitespace, then fixed-width name, then bar, then padded value
		lines = append(lines, fmt.Sprintf(" %s %s %s", name, bar, count))
	}
	return strings.Join(lines, "\n")
}
