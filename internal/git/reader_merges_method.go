package git

import (
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// GetMergeHistory runs "git log --merges --stat" to find branch merges.
func (r *Reader) GetMergeHistory() ([]domain.Stat, error) {
	// %h=hash, %s=subject, %aI=date
	output, err := r.runGit("log", "--merges", "--stat", "--format=commit:%h|%aI|%s")
	if err != nil {
		return nil, err
	}

	return parseMergeHistory(output), nil
}

// parseMergeHistory parses the interleaved commit and stat lines for merges.
func parseMergeHistory(output string) []domain.Stat {
	var stats []domain.Stat
	lines := strings.Split(strings.TrimSpace(output), "\n")

	var currentHash, currentDate, currentMessage string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "commit:") {
			parts := strings.SplitN(strings.TrimPrefix(line, "commit:"), "|", 3)
			if len(parts) >= 3 {
				currentHash, currentDate, currentMessage = parts[0], parts[1], parts[2]
			}
		} else if strings.Contains(line, "changed") && strings.Contains(line, "insertion") {
			files, added, _ := extractChurnNumbers(line) // reuse churn parser
			stats = append(stats, domain.Stat{
				RepoID:  currentHash,
				Date:    currentDate,
				Message: currentMessage,
				Value:   added, // total lines brought in from the branch
				Value3:  files, // files changed
			})
		}
	}
	return stats
}
