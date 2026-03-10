package git

import (
	"fmt" 

	"github.com/aryanwalia2003/git-stats/internal/domain" 
)

// GetRecentCommits runs "git log" to fetch the N most recent commits.
// Each commit becomes a domain.Stat with Label=author, Date=timestamp.
func (r *Reader) GetRecentCommits(limit int) ([]domain.Stat, error) {
	// --max-count=N limits how many commits we fetch
	// --format gives us a custom output: hash|author|date|message
	// The %h %an %aI %s are git placeholders:
	//   %h  = short commit hash (e.g., "a1b2c3d")
	//   %an = author name (e.g., "Aryan Walia")
	//   %aI = author date in ISO format (e.g., "2024-03-04T10:30:00+05:30")
	//   %s  = commit subject/message (e.g., "added login feature")
	maxCount := fmt.Sprintf("--max-count=%d", limit) 

	output, err := r.runGit("log", "--all", maxCount, "--format=%h|%an|%aI|%s")
	if err != nil {
		return nil, err
	}

	return parseCommitLog(output), nil 
}
