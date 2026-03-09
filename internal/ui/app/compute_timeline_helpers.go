package app

import (
	"fmt"
	"time"

	"github.com/aryanwalia2003/git-stats/internal/domain"
	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// themeValue colors a string dynamically
func themeValue(v string) string {
	return theme.ValueStyle.Render(v)
}

// formatDate returns just the YYYY-MM-DD
func formatDate(isoDate string) string {
	if len(isoDate) >= 10 {
		return isoDate[:10]
	}
	return isoDate
}

// findLongestGap finds the largest time span between any two commits.
func findLongestGap(history []domain.Stat) TimelineEvent {
	if len(history) < 2 {
		return TimelineEvent{} // need at least 2 commits to have a gap
	}

	var maxGap time.Duration
	var gapEnd time.Time

	// History is newest-first. We iterate backwards (older to newer)
	for i := len(history) - 1; i > 0; i-- {
		oldDate, _ := time.Parse(time.RFC3339, history[i].Date)
		newDate, _ := time.Parse(time.RFC3339, history[i-1].Date)

		gap := newDate.Sub(oldDate)
		if gap > maxGap {
			maxGap = gap
			gapEnd = newDate
		}
	}

	days := int(maxGap.Hours() / 24)
	if days < 7 {
		return TimelineEvent{} // Only report gaps longer than a week
	}

	return TimelineEvent{
		Emoji:       "🏖️",
		Title:       "The Great Void",
		Description: fmt.Sprintf("%d days without a single commit", days),
		Date:        formatDate(gapEnd.Format(time.RFC3339)), // date when gap ended
	}
}
