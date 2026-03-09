package app

import (
	"github.com/aryanwalia2003/git-stats/internal/domain"
	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// computeBiggestMerge finds the merge commit that brought in the most code.
func computeBiggestMerge(merges []domain.Stat) string {
	if len(merges) == 0 {
		return theme.SubtleStyle.Render("No merges found")
	}

	var biggest domain.Stat
	maxLines := -1

	for _, m := range merges {
		if m.Value > maxLines {
			maxLines = m.Value
			biggest = m
		}
	}

	return formatCommitDrilldown(biggest, biggest.Value, "lines merged")
}

// Struct to hold joining and leaving dates for a contributor
type Lifecycle struct {
	Name      string
	JoinDate  string
	LeftDate  string // will be empty if still active
	TotalDays int
}
