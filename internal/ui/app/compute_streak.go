package app

import "time"

// longestConsecutiveDays takes sorted date strings ("2026-01-01") and finds
// the longest run of consecutive calendar days.
func longestConsecutiveDays(sortedDates []string) int {
	if len(sortedDates) == 0 {
		return 0
	}
	maxStreak, current := 1, 1

	for i := 1; i < len(sortedDates); i++ {
		prev, _ := time.Parse("2006-01-02", sortedDates[i-1])
		curr, _ := time.Parse("2006-01-02", sortedDates[i])

		if curr.Sub(prev).Hours() == 24 { // exactly 1 day apart
			current++
			if current > maxStreak {
				maxStreak = current
			}
		} else {
			current = 1
		}
	}
	return maxStreak
}
