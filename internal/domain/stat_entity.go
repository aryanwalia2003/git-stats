package domain

// Stat represents a point-in-time snapshot of a metric.
type Stat struct {
	RepoID string
	Label  string
	Value  int
	Date   string
}
