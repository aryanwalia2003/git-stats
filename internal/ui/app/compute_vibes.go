package app

import (
	"time"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// computeNightOwlIndex calculates what % of commits were made between midnight and 6 AM.
func computeNightOwlIndex(commits []domain.Stat) float64 {
	if len(commits) == 0 {
		return 0
	}
	nightCount := 0
	for _, c := range commits {
		hour := extractHour(c.Date)
		if hour >= 0 && hour < 6 {
			nightCount++
		}
	}
	return float64(nightCount) / float64(len(commits)) * 100
}

// computeWeekendWarrior calculates what % of commits were on weekends.
func computeWeekendWarrior(commits []domain.Stat) float64 {
	if len(commits) == 0 {
		return 0
	}
	weekendCount := 0
	for _, c := range commits {
		day := extractWeekday(c.Date)
		if day == time.Saturday || day == time.Sunday {
			weekendCount++
		}
	}
	return float64(weekendCount) / float64(len(commits)) * 100
}
