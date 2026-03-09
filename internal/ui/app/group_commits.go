package app

import "github.com/aryanwalia2003/git-stats/internal/domain"

// groupCommitsByDate counts how many commits happened on each unique date.
// Returns parallel slices of date labels and counts (sorted by date from commits).
func groupCommitsByDate(commits []domain.Stat) ([]string, []int) {
	dateOrder := []string{}     // preserves order of appearance
	dateCounts := map[string]int{} // date -> count

	for _, c := range commits {
		date := c.Date
		if len(date) >= 10 {
			date = date[:10] // "2026-03-04T10:30:00+05:30" → "2026-03-04"
		}
		if _, exists := dateCounts[date]; !exists {
			dateOrder = append(dateOrder, date)
		}
		dateCounts[date]++
	}

	counts := make([]int, len(dateOrder))
	for i, d := range dateOrder {
		counts[i] = dateCounts[d]
	}
	return dateOrder, counts
}
