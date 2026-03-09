package app

import (
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// Unicode block characters from lowest to highest
// These create the visual "bars" in our sparkline
var sparkBlocks = []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}

// renderSparkline takes a slice of integer values and returns a sparkline string.
// Example output: "▂▄▇█▃▁▅"
func renderSparkline(values []int) string {
	if len(values) == 0 {
		return ""
	}

	// Find the max value to normalize the bars
	maxVal := 0
	for _, v := range values {
		if v > maxVal {
			maxVal = v
		}
	}
	if maxVal == 0 {
		return strings.Repeat(sparkBlocks[0], len(values))
	}

	var spark string
	for _, v := range values {
		// Scale value to 0-7 index range
		idx := (v * 7) / maxVal
		spark += theme.ValueStyle.Render(sparkBlocks[idx])
	}
	return spark
}
