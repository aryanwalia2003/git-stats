package app

import (
	"fmt"
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// findBugWar finds the day with the highest number of "fix" or "bug" keywords in commit messages.
func findBugWar(history []domain.Stat) TimelineEvent {
	counts := make(map[string]int)
	var maxDay string
	maxCount := 0

	for _, c := range history {
		msg := strings.ToLower(c.Message)
		if strings.Contains(msg, "fix") || strings.Contains(msg, "bug") {
			day := formatDate(c.Date)
			counts[day]++
			if counts[day] > maxCount {
				maxCount = counts[day]
				maxDay = day
			}
		}
	}

	if maxCount < 5 { // Only consider it a "war" if there were 5+ bugfixes in one day
		return TimelineEvent{}
	}

	return TimelineEvent{
		Emoji:       "🪲",
		Title:       "The Bug War",
		Description: fmt.Sprintf("%d bug fixes pushed in a single day", maxCount),
		Date:        maxDay,
	}
}
