package app

import (
	"time"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// extractHour parses an ISO date string and returns the hour (0-23).
func extractHour(isoDate string) int {
	t, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return -1
	}
	return t.Hour()
}

// extractWeekday parses an ISO date string and returns the day of the week.
func extractWeekday(isoDate string) time.Weekday {
	t, err := time.Parse(time.RFC3339, isoDate)
	if err != nil {
		return time.Monday
	}
	return t.Weekday()
}

// computeBusiestHour finds which hour of the day has the most commits.
func computeBusiestHour(commits []domain.Stat) (int, int) {
	hours := make([]int, 24) // 24 slots for each hour
	for _, c := range commits {
		h := extractHour(c.Date)
		if h >= 0 {
			hours[h]++
		}
	}
	bestHour, bestCount := 0, 0
	for h, count := range hours {
		if count > bestCount {
			bestHour, bestCount = h, count
		}
	}
	return bestHour, bestCount
}
