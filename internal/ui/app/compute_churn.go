package app

import (
	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// ChurnStats holds advanced code churn metrics
type ChurnStats struct {
	TotalAdded         int
	TotalDeleted       int
	TotalFiles         int
	AvgLines           int
	AvgFiles           int
	BiggestCommit      domain.Stat
	BiggestCommitFiles domain.Stat
	SmallestCommit     domain.Stat
	BiggestRefactor    domain.Stat
}

// computeChurnStats calculates advanced churn metrics from git log --shortstat output.
func computeChurnStats(churn []domain.Stat) ChurnStats {
	if len(churn) == 0 {
		return ChurnStats{}
	}

	totalA, totalD, totalF := 0, 0, 0
	maxTotal, minTotal := -1, -1
	maxRefactor, maxFiles := -1, -1
	var bigC, smallC, refactorC, bigFilesC domain.Stat

	for _, c := range churn {
		totalA += c.Value     // insertions
		totalD += c.Value2    // deletions
		totalF += c.Value3    // files changed
		totalChanges := c.Value + c.Value2

		if maxTotal == -1 || totalChanges > maxTotal {
			maxTotal = totalChanges
			bigC = c
		}
		if (minTotal == -1 || totalChanges < minTotal) && totalChanges > 0 {
			minTotal = totalChanges
			smallC = c
		}
		if c.Value2 > maxRefactor {
			maxRefactor = c.Value2
			refactorC = c
		}
		if c.Value3 > maxFiles {
			maxFiles = c.Value3
			bigFilesC = c
		}
	}

	avgLines, avgFiles := 0, 0
	if len(churn) > 0 {
		avgLines = (totalA + totalD) / len(churn)
		avgFiles = totalF / len(churn)
	}

	return ChurnStats{
		TotalAdded:         totalA,
		TotalDeleted:       totalD,
		TotalFiles:         totalF,
		AvgLines:           avgLines,
		AvgFiles:           avgFiles,
		BiggestCommit:      bigC,
		BiggestCommitFiles: bigFilesC,
		SmallestCommit:     smallC,
		BiggestRefactor:    refactorC,
	}
}
