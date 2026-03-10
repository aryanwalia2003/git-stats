package app

import (
	"fmt"
	"time"

	"github.com/aryanwalia2003/git-stats/internal/ui/theme"
)

// renderVibesPanel shows the fun behavioral stats in a single panel.
func renderVibesPanel(m Model) string {
	title := theme.PanelTitleStyle.Render("🎭 Developer Vibes")

	nightOwl := computeNightOwlIndex(m.Commits)
	weekend := computeWeekendWarrior(m.Commits)
	bestHour, _ := computeBusiestHour(m.Commits)
	bestDay, _ := computeBusiestDay(m.Commits)
	streak := computeStreak(m.Commits)
	avgMsg := computeAvgMessageLength(m.Commits)
	topEmoji, emojiCount := computeTopEmoji(m.Commits)
	caffeineActive, caffeineBurst := computeCaffeineMode(m.Commits)

	lines := fmt.Sprintf("  🌙 Night Owl:       %s\n", theme.ValueStyle.Render(fmt.Sprintf("%.0f%%", nightOwl)))
	lines += fmt.Sprintf("  ⚔️  Weekend Warrior: %s\n", theme.ValueStyle.Render(fmt.Sprintf("%.0f%%", weekend)))
	lines += fmt.Sprintf("  🗓️  Busiest:         %s at %s\n",
		theme.ValueStyle.Render(bestDay.String()[:3]),
		theme.ValueStyle.Render(fmt.Sprintf("%d:00", bestHour)))
	lines += fmt.Sprintf("  🔥 Streak:          %s days\n", theme.ValueStyle.Render(fmt.Sprintf("%d", streak)))
	lines += fmt.Sprintf("  🗯️  Avg Message:     %s chars\n", theme.ValueStyle.Render(fmt.Sprintf("%d", avgMsg)))

	if emojiCount > 0 {
		lines += fmt.Sprintf("  👑 Top Emoji:       %s x%d\n", topEmoji, emojiCount)
	} else {
		lines += fmt.Sprintf("  👑 Top Emoji:       %s\n", theme.SubtleStyle.Render("No emojis found"))
	}

	if caffeineActive {
		lines += fmt.Sprintf("  ☕ Caffeine Mode:   %s\n",
			theme.ValueStyle.Render(fmt.Sprintf("ON! %d commits in 1hr", caffeineBurst)))
	} else {
		lines += fmt.Sprintf("  ☕ Caffeine Mode:   %s\n", theme.SubtleStyle.Render("off"))
	}

	_ = time.Now // ensure time import is used
	return title + "\n" + lines
}
