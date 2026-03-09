package domain

// LocalGitReader handles extracting data from the local .git folder
type LocalGitReader interface {
	GetCurrentRepo() (*Repo, error)
	GetRecentCommits(limit int) ([]Stat, error)
	GetCommitHistory() ([]Stat, error) // gets entire commit history
	GetLocalContributors() ([]Stat, error)
	GetCodeChurn() ([]Stat, error)
}
