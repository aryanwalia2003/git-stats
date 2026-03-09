package app

import "github.com/aryanwalia2003/git-stats/internal/domain"

// StatsLoadedMsg is sent when all git stats have been fetched successfully.
type StatsLoadedMsg struct {
	Contributors []domain.Stat // top committers ranked by count
	Commits      []domain.Stat // recent commits with author + date
	Churn        []domain.Stat // lines added/deleted per commit
	History      []domain.Stat // full commit history for timeline
}

// StatsErrorMsg is sent when fetching stats fails.
type StatsErrorMsg struct {
	Err error
}
