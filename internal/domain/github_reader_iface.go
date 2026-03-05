package domain

// GitHubReader fetches remote repo statistics
type GitHubReader interface {
	GetTraffic(repoOwner, repoName string) ([]Stat, error) // gets traffic
	GetOpenIssues(repoOwner, repoName string) ([]Stat, error) // gets open issues
	GetStargazers(repoOwner, repoName string) ([]Stat, error) // gets stargazers
}
