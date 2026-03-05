package domain

// Repo represents the singular local repository context.
type Repo struct {
	LocalPath    string
	RemoteOrigin string
	Name         string
	Owner        string
	Branch       string
	Description  string // Pulled from GitHub, not local
	StarCount    int    // Pulled from GitHub
	ForkCount    int    // Pulled from GitHub
}
