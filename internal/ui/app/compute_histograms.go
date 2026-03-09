package app

import "github.com/aryanwalia2003/git-stats/internal/domain"

// computeHourHistogram counts commits per hour of the day (0-23).
func computeHourHistogram(commits []domain.Stat) []int {
	hours := make([]int, 24)
	for _, c := range commits {
		h := extractHour(c.Date)
		if h >= 0 {
			hours[h]++
		}
	}
	return hours
}

// computeDayHistogram counts commits per day of the week (Sun=0 to Sat=6).
func computeDayHistogram(commits []domain.Stat) []int {
	days := make([]int, 7)
	for _, c := range commits {
		d := extractWeekday(c.Date)
		days[d]++
	}
	return days
}
