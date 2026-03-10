package git

import "github.com/aryanwalia2003/git-stats/internal/domain"

// GetCommitHistory runs "git log" without a limit to fetch ALL commits.
// This is used to build the rich timeline (start date, big gaps, etc).
func (r *Reader) GetCommitHistory() ([]domain.Stat, error) {
	// No --max-count here, we want everything.
	// %h=hash, %an=author, %aI=date, %s=subject
	output, err := r.runGit("log", "--all", "--format=%h|%an|%aI|%s")
	if err != nil {
		return nil, err
	}

	return parseCommitLog(output), nil // reuse the existing parser
}
