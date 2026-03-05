package git

import "github.com/aryanwalia2003/git-stats/internal/domain"

// GetCodeChurn runs "git log --shortstat" to find how many lines
// were added and deleted across the repo's recent history.
func (r *Reader) GetCodeChurn() ([]domain.Stat, error) {
	// --shortstat gives us a summary line after each commit like:
	//   "3 files changed, 42 insertions(+), 10 deletions(-)"
	// --format="" suppresses the normal commit info so we only get stats
	output, err := r.runGit("log", "--shortstat", "--format=commit:%h|%an")
	if err != nil {
		return nil, err
	}

	return parseChurnOutput(output), nil // delegate to parse_churn.go
}
