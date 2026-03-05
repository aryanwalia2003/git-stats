
package git

import (
	"strconv"
	"strings"
	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// parseContributorsStats converts git shortlog output into domain.Stat
// Input example line: "   142  Aryan Walia"
func parseContributorsStats(output string) []domain.Stat {
	var stats []domain.Stat
	lines := strings.Split(strings.TrimSpace(output), "\n")
	
	for _, line := range lines {
		parts := strings.Fields(line) // Splits by any whitespace
		if len(parts) >= 2 {
			count, _ := strconv.Atoi(parts[0])
			author := strings.Join(parts[1:], " ")
			
			stats = append(stats, domain.Stat{
				Label: author,
				Value: count,
			})
		}
	}
	return stats
}
