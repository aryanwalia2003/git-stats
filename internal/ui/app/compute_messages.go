package app

import (
	"unicode/utf8"

	"github.com/aryanwalia2003/git-stats/internal/domain"
)

// computeAvgMessageLength calculates the average commit message length.
func computeAvgMessageLength(commits []domain.Stat) int {
	if len(commits) == 0 {
		return 0
	}
	total := 0
	for _, c := range commits {
		total += utf8.RuneCountInString(c.Message)
	}
	return total / len(commits)
}

// computeTopEmoji finds the most frequently used emoji in commit messages.
func computeTopEmoji(commits []domain.Stat) (string, int) {
	emojiCounts := map[string]int{}
	for _, c := range commits {
		for _, r := range c.Message {
			if isEmoji(r) {
				emojiCounts[string(r)]++
			}
		}
	}
	bestEmoji, bestCount := "🤷", 0 // default if no emoji found
	for emoji, count := range emojiCounts {
		if count > bestCount {
			bestEmoji, bestCount = emoji, count
		}
	}
	return bestEmoji, bestCount
}
