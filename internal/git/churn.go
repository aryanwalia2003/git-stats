package git

import "github.com/aryanwalia2003/git-stats/internal/domain"

// GetCodeChurn calculates lines added/removed across recent history.
func (r *Reader) GetCodeChurn() ([]domain.Stat, error) {
	// TODO: run git log --shortstat
	return nil, nil
}
