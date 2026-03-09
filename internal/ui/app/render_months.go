package app

import (
	"fmt"
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
	"github.com/charmbracelet/lipgloss"
)

// renderMonthBlocks renders the 6 distinct month blocks side-by-side.
func renderMonthBlocks(blocks []MonthBlock) string {
	if len(blocks) == 0 {
		return ""
	}

	var renderedBlocks []string
	blockChar := "■ "

	for _, block := range blocks {
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("\n  %s\n", theme.LabelStyle.Render(block.Name)))
		
		// Render the max 31 days into ~7 rows of 5 columns to form a neat "block" (5x7 max)
		// LeetCode does continuous weeks, but user wants literal monthly squares.
		// A 5-column (weeks) x 7-row (days) grid fits 35 days perfectly.
		
		for row := 0; row < 7; row++ {
			sb.WriteString("  ")
			for col := 0; col < 5; col++ {
				dayIndex := (col * 7) + row
				
				if dayIndex >= block.Total {
					// empty space if month has 30/28 days
					sb.WriteString(theme.SubtleStyle.Render("  "))
					continue
				}

				count := block.Days[dayIndex]
				colorIdx := count
				if count >= 4 { colorIdx = 4 }

				style := lipgloss.NewStyle().Foreground(theme.CalendarColors[colorIdx])
				sb.WriteString(style.Render(blockChar))
			}
			sb.WriteString("\n")
		}
		renderedBlocks = append(renderedBlocks, sb.String())
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, renderedBlocks...)
}
