package git

import (
	"strings" // strings has helpers to split text by delimiters

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// parseCommitLog converts the custom git log output into domain.Stat slices.
// Each input line looks like: "a1b2c3d|Aryan Walia|2024-03-04T10:30:00+05:30|added login"
func parseCommitLog(output string) []domain.Stat {
	var stats []domain.Stat                                // empty slice to collect results
	lines := strings.Split(strings.TrimSpace(output), "\n") // split output into lines

	for _, line := range lines {
		// Split each line by "|" into 4 parts: [hash, author, date, message]
		parts := strings.SplitN(line, "|", 4) // SplitN with 4 means max 4 pieces
		if len(parts) < 4 {
			continue // skip malformed lines (safety check)
		}

		stats = append(stats, domain.Stat{
			RepoID: parts[0], // commit hash — unique ID for this commit
			Label:  parts[1], // author name
			Date:   parts[2], // ISO timestamp
			// parts[3] is the commit message — we could store it later
		})
	}
	return stats
}
