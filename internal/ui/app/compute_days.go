package app

import (
	"sort"
	"time"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// computeBusiestDay finds which day of the week has the most commits.
func computeBusiestDay(commits []domain.Stat) (time.Weekday, int) {
	days := make([]int, 7) // Sun=0 through Sat=6
	for _, c := range commits {
		day := extractWeekday(c.Date)
		days[day]++
	}
	bestDay, bestCount := time.Sunday, 0
	for d, count := range days {
		if count > bestCount {
			bestDay, bestCount = time.Weekday(d), count
		}
	}
	return bestDay, bestCount
}

// computeStreak finds longest consecutive days with at least one commit.
func computeStreak(commits []domain.Stat) int {
	if len(commits) == 0 {
		return 0
	}
	// Collect unique dates
	dateSet := map[string]bool{}
	for _, c := range commits {
		if len(c.Date) >= 10 {
			dateSet[c.Date[:10]] = true
		}
	}
	// Sort dates
	dates := make([]string, 0, len(dateSet))
	for d := range dateSet {
		dates = append(dates, d)
	}
	sort.Strings(dates)

	return longestConsecutiveDays(dates)
}
