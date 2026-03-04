package domain

// LocalGitReader handles extracting data from the local .git folder
type LocalGitReader interface {
	GetCurrentRepo() (*Repo, error) // gets current repo
	GetRecentCommits(limit int) ([]Stat, error) // gets recent commits
	GetLocalContributors() ([]Stat, error) // gets local contributors (local contributors are those who have contributed to the repo)
	GetCodeChurn() ([]Stat, error) // gets code churn (code churn is the amount of code that has been added or removed from the repo)
}
