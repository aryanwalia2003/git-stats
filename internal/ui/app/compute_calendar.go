package app

import (
	"time"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// computeCalendarMatrix builds a 7-day row by N-week column matrix of commit counts.
// Matrix[0] is Sunday, Matrix[6] is Saturday.
func computeCalendarMatrix(history []domain.Stat, weeks int) [][]int {
	matrix := make([][]int, 7)
	for i := range matrix {
		matrix[i] = make([]int, weeks)
	}

	if len(history) == 0 {
		return matrix
	}

	// We want the most recent 'weeks' ending today.
	now := time.Now()
	// Find the Sunday of the oldest week we care about
	startOfCurrentWeek := now.AddDate(0, 0, -int(now.Weekday()))
	startOfGraph := startOfCurrentWeek.AddDate(0, 0, -(weeks-1)*7)

	for _, stat := range history {
		t, err := time.Parse(time.RFC3339, stat.Date)
		if err != nil || t.Before(startOfGraph) {
			continue // skip commits older than our graph
		}

		daysSinceStart := int(t.Sub(startOfGraph).Hours() / 24)
		if daysSinceStart < 0 || daysSinceStart >= weeks*7 {
			continue
		}

		col := daysSinceStart / 7
		row := int(t.Weekday())
		matrix[row][col]++
	}

	return matrix
}
