package app

import (
	"time"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// MonthBlock represents a single month in our 6-month calendar view.
type MonthBlock struct {
	Name  string // e.g., "Oct", "Nov"
	Days  []int  // Array representing days 1 to 31 (indexed 0-30)
	Total int    // Total days in this month
}

// computeMonthlyBlocks groups the history into explicit Calendar Months (dynamically sized based on repo age up to maxMonths).
func computeMonthlyBlocks(history []domain.Stat, maxMonths int) []MonthBlock {
	if len(history) == 0 {
		return []MonthBlock{}
	}

	now := time.Now()

	// History is newest-first. The oldest commit is at the end.
	oldestStat := history[len(history)-1]
	oldestTime, _ := time.Parse(time.RFC3339, oldestStat.Date)
	if oldestTime.IsZero() {
		oldestTime = now
	}

	monthsSinceOldest := int(now.Month()-oldestTime.Month()) + 12*(now.Year()-oldestTime.Year())
	numMonths := monthsSinceOldest + 1

	if numMonths > maxMonths {
		numMonths = maxMonths
	}
	if numMonths <= 0 {
		numMonths = 1
	}

	blocks := make([]MonthBlock, numMonths)

	// Initialize the empty month blocks (Current month is at the end, index numMonths-1)
	for i := 0; i < numMonths; i++ {
		// Go back (numMonths - 1 - i) months
		targetDate := now.AddDate(0, -(numMonths - 1 - i), 0)
		
		// Find days in this target month
		firstOfMonth := time.Date(targetDate.Year(), targetDate.Month(), 1, 0, 0, 0, 0, time.UTC)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
		
		blocks[i] = MonthBlock{
			Name:  targetDate.Month().String()[:3],    // "Jan", "Feb"
			Days:  make([]int, lastOfMonth.Day()),     // e.g., slice of 31 for Jan
			Total: lastOfMonth.Day(),
		}
	}

	// Safety check
	if len(history) == 0 {
		return blocks
	}

	for _, stat := range history {
		t, err := time.Parse(time.RFC3339, stat.Date)
		if err != nil {
			continue
		}

		// Calculate how many months ago this commit was compared to *current month*
		monthsAgo := int(now.Month()-t.Month()) + 12*(now.Year()-t.Year())
		
		// If the commit belongs in one of our tracked blocks
		if monthsAgo >= 0 && monthsAgo < numMonths {
			blockIndex := numMonths - 1 - monthsAgo
			dayIndex := t.Day() - 1 // zero-indexed
			
			if blockIndex >= 0 && dayIndex >= 0 && dayIndex < blocks[blockIndex].Total {
				blocks[blockIndex].Days[dayIndex]++
			}
		}
	}

	return blocks
}
