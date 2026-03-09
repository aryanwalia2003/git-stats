package app

import (
	"time"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// computeLifecycles analyzes when people joined and if they took a long sabbatical/left.
// History is newest-first.
func computeLifecycles(history []domain.Stat) []Lifecycle {
	lastSeen := make(map[string]time.Time)
	firstSeen := make(map[string]time.Time)

	for _, c := range history {
		date, _ := time.Parse(time.RFC3339, c.Date)
		author := c.Label

		// Since history is newest-to-oldest, the first time we see them loop-wise is their LAST commit.
		if _, exists := lastSeen[author]; !exists {
			lastSeen[author] = date
		}
		// The last time we see them loop-wise will be their FIRST commit (handled by constantly overwriting).
		firstSeen[author] = date
	}

	var lifecycles []Lifecycle
	now := time.Now()

	for author, joinDate := range firstSeen {
		lastDate := lastSeen[author]
		daysSinceLastCommit := int(now.Sub(lastDate).Hours() / 24)
		totalDaysActive := int(lastDate.Sub(joinDate).Hours() / 24)

		left := ""
		if daysSinceLastCommit > 60 {
			left = formatDate(lastDate.Format(time.RFC3339))
		}

		lifecycles = append(lifecycles, Lifecycle{
			Name:      author,
			JoinDate:  formatDate(joinDate.Format(time.RFC3339)),
			LeftDate:  left,
			TotalDays: totalDaysActive,
		})
	}

	return lifecycles
}
