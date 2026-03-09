package app

import (
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
	"github.com/charmbracelet/lipgloss"
)

// renderCalendar builds a 7xN LeetCode-style contribution heatmap.
func renderCalendar(matrix [][]int) string {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return ""
	}

	weeks := len(matrix[0])
	var sb strings.Builder

	// The block character used for the grid squares
	block := "■ "

	for row := 0; row < 7; row++ {
		// Row prefixes: show days
		dayLabel := "    "
		if row == 1 {
			dayLabel = "Mon "
		} else if row == 3 {
			dayLabel = "Wed "
		} else if row == 5 {
			dayLabel = "Fri "
		}
		sb.WriteString(theme.SubtleStyle.Render(dayLabel))

		// Render the columns (weeks) for this row (day)
		for col := 0; col < weeks; col++ {
			count := matrix[row][col]
			colorIdx := count

			// Max cap at 4+ commits
			if count >= 4 {
				colorIdx = 4
			}

			// Style this specific block
			style := lipgloss.NewStyle().Foreground(theme.CalendarColors[colorIdx])
			sb.WriteString(style.Render(block))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
