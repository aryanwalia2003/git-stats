package git

import "github.com/aryanwalia2003/git-stats/internal/domain"

// GetCurrentRepo reads .git/config to determine the current repository info.
func (r *Reader) GetCurrentRepo() (*domain.Repo, error) {
	// TODO: parse .git/config for remote origin URL
	return nil, nil
}
