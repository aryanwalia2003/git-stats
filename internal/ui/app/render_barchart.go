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
		bar := theme.ValueStyle.Render(strings.Repeat("█", barLen))
		count := theme.SubtleStyle.Render(fmt.Sprintf(" %d", values[i]))
		name := theme.LabelStyle.Render(fmt.Sprintf("%-16s", label))
		lines = append(lines, fmt.Sprintf("  %s %s%s", name, bar, count))
	}
	return strings.Join(lines, "\n")
}
