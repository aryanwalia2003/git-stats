package git

import (
	"strconv" // strconv.Atoi converts strings like "42" to int 42
	"strings" // strings has Split, Contains, Fields etc.

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// parseChurnOutput reads interleaved commit + shortstat lines.
// Input looks like:
//   commit:a1b2c3d|Aryan Walia
//    3 files changed, 42 insertions(+), 10 deletions(-)
func parseChurnOutput(output string) []domain.Stat {
	var stats []domain.Stat                                 // result slice
	lines := strings.Split(strings.TrimSpace(output), "\n") // split into lines

	var currentAuthor string // tracks who made the current commit

	for _, line := range lines {
		line = strings.TrimSpace(line) // remove leading/trailing spaces

		if strings.HasPrefix(line, "commit:") {
			// This is a commit header line: "commit:a1b2c3d|Aryan"
			parts := strings.SplitN(strings.TrimPrefix(line, "commit:"), "|", 2)
			if len(parts) == 2 {
				currentAuthor = parts[1] // save author for the next stat line
			}
		} else if strings.Contains(line, "changed") {
			// This is a shortstat line: "3 files changed, 42 insertions(+)"
			added, deleted := extractChurnNumbers(line)
			stats = append(stats, domain.Stat{
				Label: currentAuthor, // who made this churn
				Value: added - deleted, // net change (positive = growth)
			})
		}
	}
	return stats
}

// extractChurnNumbers pulls the insertions and deletions counts from a shortstat line.
// Input: "3 files changed, 42 insertions(+), 10 deletions(-)"
// Returns: added=42, deleted=10
func extractChurnNumbers(line string) (added int, deleted int) {
	fields := strings.Fields(line) // split by whitespace into words

	for i, word := range fields {
		// Look for the word right before "insertions(+)" — that's the added count
		if strings.HasPrefix(word, "insertion") && i > 0 {
			added, _ = strconv.Atoi(fields[i-1]) // "42" -> 42
		}
		// Look for the word right before "deletions(-)" — that's the deleted count
		if strings.HasPrefix(word, "deletion") && i > 0 {
			deleted, _ = strconv.Atoi(fields[i-1]) // "10" -> 10
		}
	}
	return added, deleted
}
