package git

import "github.com/aryanwalia2003/git-stats/internal/domain" // our own domain entities

// GetCurrentRepo reads the .git/config to find the remote origin URL,
// then extracts the owner and repo name from it.
func (r*Reader) GetCurrentRepo() (*domain.Repo, error) {
	// This runs: git config --get remote.origin.url
	// Which returns something like: https://github.com/aryanwalia2003/git-stats.git
	originURL, err := r.runGit("config", "--get", "remote.origin.url")
	if err != nil {
		return nil, err // if no remote is set, return the error
	}

	// This runs: git branch --show-current
	// Which returns the current branch name like: main
	branch, err := r.runGit("branch", "--show-current")
	if err != nil {
		return nil, err
	}

	// Parse "aryanwalia2003" and "git-stats" from the URL
	owner, name := parseOriginURL(originURL)

	// Build and return the Repo struct with all local info filled in
	return &domain.Repo{
		LocalPath:    r.targetDir, // the folder path we were given
		RemoteOrigin: originURL,   // full remote URL
		Name:         name,        // just the repo name
		Owner:        owner,       // the GitHub username
		Branch:       branch,      // current branch
	}, nil
}
