package app

import (
	"sort"
	"time"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// computeCaffeineMode checks if there's a burst of 5+ commits within any 1-hour window.
func computeCaffeineMode(commits []domain.Stat) (bool, int) {
	if len(commits) < 5 {
		return false, 0
	}

	// Parse all timestamps and sort them
	times := make([]time.Time, 0, len(commits))
	for _, c := range commits {
		t, err := time.Parse(time.RFC3339, c.Date)
		if err == nil {
			times = append(times, t)
		}
	}
	sort.Slice(times, func(i, j int) bool { return times[i].Before(times[j]) })

	// Sliding window: check if 5+ commits are within 1 hour
	maxBurst := 0
	for i := 0; i < len(times); i++ {
		count := 1
		for j := i + 1; j < len(times); j++ {
			if times[j].Sub(times[i]) <= time.Hour {
				count++
			} else {
				break
			}
		}
		if count > maxBurst {
			maxBurst = count
		}
	}
	return maxBurst >= 5, maxBurst
}
