package app

import (
	"fmt"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// findBusiestDayInHistory groups all commits by YYYY-MM-DD to find the absolute peak activity day.
func findBusiestDayInHistory(history []domain.Stat) TimelineEvent {
	counts := make(map[string]int)
	var maxDay string
	maxCount := 0

	for _, c := range history {
		day := formatDate(c.Date)
		counts[day]++
		if counts[day] > maxCount {
			maxCount = counts[day]
			maxDay = day
		}
	}

	if maxCount < 10 {
		return TimelineEvent{} // Ignore if max is literally like 3 commits.
	}

	return TimelineEvent{
		Emoji:       "🐝",
		Title:       "Busy Bee",
		Description: fmt.Sprintf("%d commits pushed in a single day", maxCount),
		Date:        maxDay,
	}
}
