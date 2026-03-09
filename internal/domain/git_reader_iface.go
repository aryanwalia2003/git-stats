package domain

// LocalGitReader handles extracting data from the local .git folder
type LocalGitReader interface {
	GetCurrentRepo() (*Repo, error)
	GetRecentCommits(limit int) ([]Stat, error)
	GetCommitHistory() ([]Stat, error) // gets entire commit history
	GetMergeHistory() ([]Stat, error)  // gets merge/PR history for branch stats
	GetLocalContributors() ([]Stat, error)
	GetCodeChurn() ([]Stat, error)
	GetFileFrequencies() ([]Stat, error) // counts how many times files were modified
}
