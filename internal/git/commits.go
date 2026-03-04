package git

import "github.com/aryanwalia2003/git-stats/internal/domain"

// GetRecentCommits parses git log for the N most recent commits.
func (r *Reader) GetRecentCommits(limit int) ([]domain.Stat, error) {
	// TODO: run git log --max-count=limit
	return nil, nil
}
