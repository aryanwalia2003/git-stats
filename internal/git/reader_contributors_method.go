package git

import "github.com/aryanwalia2003/git-stats/internal/domain"

// GetLocalContributors runs "git shortlog" to get a ranked list of contributors.
// -s = suppress commit messages (only show counts)
// -n = sort by number of commits (highest first)
// --all = include all branches, not just the current one
func (r *Reader) GetLocalContributors() ([]domain.Stat, error) {
	output, err := r.runGit("shortlog", "-sn", "--all") // uses shared runGit helper
	if err != nil {
		return nil, err
	}

	return parseContributorsStats(output), nil // delegate to parse.go
}
