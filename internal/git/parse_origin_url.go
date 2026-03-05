package git

import "strings"

// parseOriginURL extracts owner and repo name from a git remote URL.
// Supports both formats:
//   https://github.com/aryanwalia2003/git-stats.git
//   git@github.com:aryanwalia2003/git-stats.git
func parseOriginURL(rawURL string) (owner string, name string) {
	// Remove trailing .git if present
	// "git-stats.git" becomes "git-stats"
	rawURL = strings.TrimSuffix(rawURL, ".git")

	// Try HTTPS format first: split by "/"
	// https://github.com/aryanwalia2003/git-stats -> parts = [..., "aryanwalia2003", "git-stats"]
	parts := strings.Split(rawURL, "/")
	if len(parts) >= 2 {
		return parts[len(parts)-2], parts[len(parts)-1]
	}

	// Fallback for SSH format: git@github.com:aryanwalia2003/git-stats
	// Split by ":" first, then by "/"
	if colonParts := strings.SplitN(rawURL, ":", 2); len(colonParts) == 2 {
		sshParts := strings.Split(colonParts[1], "/")
		if len(sshParts) == 2 {
			return sshParts[0], sshParts[1]
		}
	}

	return "", "" // couldn't parse — will be empty
}
