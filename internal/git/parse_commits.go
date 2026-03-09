package git

import (
	"strings" // strings has helpers to split text by delimiters

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// parseCommitLog converts the custom git log output into domain.Stat slices.
// Each input line looks like: "a1b2c3d|Aryan Walia|2024-03-04T10:30:00+05:30|added login"
func parseCommitLog(output string) []domain.Stat {
	var stats []domain.Stat                                
	lines := strings.Split(strings.TrimSpace(output), "\n") 

	for _, line := range lines {
		// Split each line by "|" into 4 parts: [hash, author, date, message]
		parts := strings.SplitN(line, "|", 4) 
		if len(parts) < 4 {
			continue 
		}

		stats = append(stats, domain.Stat{
			RepoID:  parts[0],
			Label:   parts[1],
			Date:    parts[2],
			Message: parts[3],
		})
	}
	return stats
}
