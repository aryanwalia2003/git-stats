package git

import (
	"bytes"
	"os/exec"
	"sort"
	"strings"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// GetFileFrequencies runs `git log --name-only` to count how many times each file was modified.
func (r *Reader) GetFileFrequencies() ([]domain.Stat, error) {
	cmd := exec.Command("git", "log", "--name-only", "--pretty=format:")
	cmd.Dir = r.targetDir

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	counts := make(map[string]int)
	lines := strings.Split(out.String(), "\n")
	
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue // skip empty lines between commits
		}
		counts[trimmed]++
	}

	var stats []domain.Stat
	for file, count := range counts {
		stats = append(stats, domain.Stat{
			Label: file,
			Value: count,
		})
	}

	// Sort by count descending
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].Value > stats[j].Value
	})

	return stats, nil
}
