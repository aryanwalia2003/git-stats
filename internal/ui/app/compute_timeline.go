package app

import (
	"fmt"
	"sort"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// TimelineEvent represents a major milestone or fun fact in the repo's history.
type TimelineEvent struct {
	Emoji       string
	Title       string
	Description string
	Date        string // when it happened
}

// computeTimeline analyzes the full commit history to find key events.
// Assumes history is sorted newest-to-oldest (default git log order).
func computeTimeline(history []domain.Stat) []TimelineEvent {
	if len(history) == 0 {
		return nil
	}

	var events []TimelineEvent

	// Git log puts newest first, so the last element is the oldest commit
	firstCommit := history[len(history)-1]
	events = append(events, TimelineEvent{
		Emoji:       "🌱",
		Title:       "The Genesis",
		Description: fmt.Sprintf("First commit by %s", themeValue(firstCommit.Label)),
		Date:        formatDate(firstCommit.Date),
	})

	// Find longest gap and busiest day
	longestGap := findLongestGap(history)
	if longestGap.Title != "" {
		events = append(events, longestGap)
	}

	busiestDay := findBusiestDayInHistory(history)
	if busiestDay.Title != "" {
		events = append(events, busiestDay)
	}

	bugWar := findBugWar(history)
	if bugWar.Title != "" {
		events = append(events, bugWar)
	}

	// Sort events chronologically (oldest to newest)
	sort.Slice(events, func(i, j int) bool {
		return events[i].Date < events[j].Date
	})

	return events
}
