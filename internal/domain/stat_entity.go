package domain

// Stat represents a point-in-time snapshot of a metric.
type Stat struct {
	RepoID  string // e.g., commit hash
	Label   string
	Value   int    // e.g., insertions
	Value2  int    // e.g., deletions
	Value3  int    // e.g., files changed
	Date    string
	Message string // commit message (for Talker/Emoji analysis)
}
